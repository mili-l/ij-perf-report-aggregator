package main

import (
  "context"
  "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
  "github.com/develar/errors"
  "golang.org/x/tools/container/intsets"
  "strconv"
  "time"
)

type ReportExistenceChecker struct {
  ids intsets.Sparse
}

func (t *ReportExistenceChecker) reset(taskContext context.Context, dbName string, tableName string, db driver.Conn, since time.Time) error {
  t.ids.Clear()

  var rows driver.Rows
  var err error
  if dbName == "ij" {
    // don't filter by machine - product is enough to reduce set
    query := "select tc_build_id from report where generated_time > $1 order by tc_build_id"
    rows, err = db.Query(taskContext, query, since)
  } else {
    table := "report"
    if tableName != "" {
      table = tableName
    }
    query := "select tc_build_id from " + table + " where generated_time > " + strconv.FormatInt(since.Unix(), 10) + " order by tc_build_id"
    rows, err = db.Query(taskContext, query, since)
  }

  if err != nil {
    return errors.WithStack(err)
  }

  for rows.Next() {
    // clickhouse requires explicit type
    var id uint32
    err = rows.Scan(&id)
    if err != nil {
      return errors.WithStack(err)
    }

    t.ids.Insert(int(id))
  }

  if rows.Err() != nil {
    return errors.WithStack(rows.Err())
  }
  return nil
}

func (t *ReportExistenceChecker) has(id int) bool {
  return t.ids.Has(id)
}
