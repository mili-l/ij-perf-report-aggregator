// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/server/json.qtpl:1
package server

//line pkg/server/json.qtpl:1
import "github.com/bvinc/go-sqlite-lite/sqlite3"

//line pkg/server/json.qtpl:2
import "github.com/develar/errors"

//line pkg/server/json.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/server/json.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/server/json.qtpl:5
func streamhttpError(qw422016 *qt422016.Writer, error *HttpError) {
//line pkg/server/json.qtpl:5
	qw422016.N().S(`{"error":`)
//line pkg/server/json.qtpl:7
	qw422016.N().Q(error.Message)
//line pkg/server/json.qtpl:7
	qw422016.N().S(`}`)
//line pkg/server/json.qtpl:9
}

//line pkg/server/json.qtpl:9
func writehttpError(qq422016 qtio422016.Writer, error *HttpError) {
//line pkg/server/json.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:9
	streamhttpError(qw422016, error)
//line pkg/server/json.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:9
}

//line pkg/server/json.qtpl:9
func httpError(error *HttpError) string {
//line pkg/server/json.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:9
	writehttpError(qb422016, error)
//line pkg/server/json.qtpl:9
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:9
	return qs422016
//line pkg/server/json.qtpl:9
}

//line pkg/server/json.qtpl:11
func StreamGroupedMetricList(qw422016 *qt422016.Writer, list []MedianResult) {
//line pkg/server/json.qtpl:11
	qw422016.N().S(`{"groupNames": [`)
//line pkg/server/json.qtpl:14
	if len(list) > 0 {
//line pkg/server/json.qtpl:15
		for i, value := range list[0].buildToValue {
//line pkg/server/json.qtpl:16
			if i != 0 {
//line pkg/server/json.qtpl:16
				qw422016.N().S(`,`)
//line pkg/server/json.qtpl:16
			}
//line pkg/server/json.qtpl:16
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:17
			qw422016.N().D(value.buildC1)
//line pkg/server/json.qtpl:17
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:18
		}
//line pkg/server/json.qtpl:19
	}
//line pkg/server/json.qtpl:19
	qw422016.N().S(`],"data": [`)
//line pkg/server/json.qtpl:22
	for i, item := range list {
//line pkg/server/json.qtpl:23
		if i != 0 {
//line pkg/server/json.qtpl:23
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:23
		}
//line pkg/server/json.qtpl:23
		qw422016.N().S(`{`)
//line pkg/server/json.qtpl:25
		for _, value := range item.buildToValue {
//line pkg/server/json.qtpl:25
			qw422016.N().S(`"`)
//line pkg/server/json.qtpl:26
			qw422016.N().D(value.buildC1)
//line pkg/server/json.qtpl:26
			qw422016.N().S(`":`)
//line pkg/server/json.qtpl:26
			qw422016.N().D(value.value)
//line pkg/server/json.qtpl:26
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:27
		}
//line pkg/server/json.qtpl:30
		qw422016.N().S(`"metric":`)
//line pkg/server/json.qtpl:31
		qw422016.N().Q(item.metricName)
//line pkg/server/json.qtpl:31
		qw422016.N().S(`}`)
//line pkg/server/json.qtpl:33
	}
//line pkg/server/json.qtpl:33
	qw422016.N().S(`]}`)
//line pkg/server/json.qtpl:36
}

//line pkg/server/json.qtpl:36
func WriteGroupedMetricList(qq422016 qtio422016.Writer, list []MedianResult) {
//line pkg/server/json.qtpl:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:36
	StreamGroupedMetricList(qw422016, list)
//line pkg/server/json.qtpl:36
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:36
}

//line pkg/server/json.qtpl:36
func GroupedMetricList(list []MedianResult) string {
//line pkg/server/json.qtpl:36
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:36
	WriteGroupedMetricList(qb422016, list)
//line pkg/server/json.qtpl:36
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:36
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:36
	return qs422016
//line pkg/server/json.qtpl:36
}

//line pkg/server/json.qtpl:38
func StreamInfo(qw422016 *qt422016.Writer, productNames []string, essentialMetricNames []string, instantMetricNames []string, statement *sqlite3.Stmt, errRef *error) {
//line pkg/server/json.qtpl:38
	qw422016.N().S(`{"productNames":`)
//line pkg/server/json.qtpl:40
	streamsafeStringList(qw422016, productNames)
//line pkg/server/json.qtpl:40
	qw422016.N().S(`,"durationMetricNames": [`)
//line pkg/server/json.qtpl:42
	for _, name := range essentialMetricNames {
//line pkg/server/json.qtpl:42
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:43
		qw422016.N().S(name)
//line pkg/server/json.qtpl:43
		qw422016.N().S(`",`)
//line pkg/server/json.qtpl:44
	}
//line pkg/server/json.qtpl:44
	qw422016.N().S(`"moduleLoading"],"instantMetricNames": [`)
//line pkg/server/json.qtpl:48
	for i, name := range instantMetricNames {
//line pkg/server/json.qtpl:49
		if i != 0 {
//line pkg/server/json.qtpl:49
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:49
		}
//line pkg/server/json.qtpl:49
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:50
		qw422016.N().S(name)
//line pkg/server/json.qtpl:50
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:51
	}
//line pkg/server/json.qtpl:51
	qw422016.N().S(`],"productToMachine": {`)
//line pkg/server/json.qtpl:54
	for i, product := range productNames {
//line pkg/server/json.qtpl:55
		if i != 0 {
//line pkg/server/json.qtpl:55
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:57
			statement.Reset()

//line pkg/server/json.qtpl:58
		}
//line pkg/server/json.qtpl:61
		*errRef = errors.WithStack(statement.BindString(product, 0))
		if *errRef != nil {
			return
		}

//line pkg/server/json.qtpl:65
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:67
		qw422016.N().S(product)
//line pkg/server/json.qtpl:67
		qw422016.N().S(`":`)
//line pkg/server/json.qtpl:67
		streamwriteMachineInfoList(qw422016, statement, errRef)
//line pkg/server/json.qtpl:68
	}
//line pkg/server/json.qtpl:68
	qw422016.N().S(`}}`)
//line pkg/server/json.qtpl:71
}

//line pkg/server/json.qtpl:71
func WriteInfo(qq422016 qtio422016.Writer, productNames []string, essentialMetricNames []string, instantMetricNames []string, statement *sqlite3.Stmt, errRef *error) {
//line pkg/server/json.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:71
	StreamInfo(qw422016, productNames, essentialMetricNames, instantMetricNames, statement, errRef)
//line pkg/server/json.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:71
}

//line pkg/server/json.qtpl:71
func Info(productNames []string, essentialMetricNames []string, instantMetricNames []string, statement *sqlite3.Stmt, errRef *error) string {
//line pkg/server/json.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:71
	WriteInfo(qb422016, productNames, essentialMetricNames, instantMetricNames, statement, errRef)
//line pkg/server/json.qtpl:71
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:71
	return qs422016
//line pkg/server/json.qtpl:71
}

//line pkg/server/json.qtpl:73
func streamsafeStringList(qw422016 *qt422016.Writer, list []string) {
//line pkg/server/json.qtpl:73
	qw422016.N().S(`[`)
//line pkg/server/json.qtpl:75
	for i, v := range list {
//line pkg/server/json.qtpl:76
		if i != 0 {
//line pkg/server/json.qtpl:76
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:76
		}
//line pkg/server/json.qtpl:76
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:77
		qw422016.N().S(v)
//line pkg/server/json.qtpl:77
		qw422016.N().S(`"`)
//line pkg/server/json.qtpl:78
	}
//line pkg/server/json.qtpl:78
	qw422016.N().S(`]`)
//line pkg/server/json.qtpl:80
}

//line pkg/server/json.qtpl:80
func writesafeStringList(qq422016 qtio422016.Writer, list []string) {
//line pkg/server/json.qtpl:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:80
	streamsafeStringList(qw422016, list)
//line pkg/server/json.qtpl:80
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:80
}

//line pkg/server/json.qtpl:80
func safeStringList(list []string) string {
//line pkg/server/json.qtpl:80
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:80
	writesafeStringList(qb422016, list)
//line pkg/server/json.qtpl:80
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:80
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:80
	return qs422016
//line pkg/server/json.qtpl:80
}

//line pkg/server/json.qtpl:82
func streamwriteMachineInfoList(qw422016 *qt422016.Writer, statement *sqlite3.Stmt, errRef *error) {
//line pkg/server/json.qtpl:82
	qw422016.N().S(`[`)
//line pkg/server/json.qtpl:84
	isFirstRow := true

//line pkg/server/json.qtpl:85
	for {
//line pkg/server/json.qtpl:87
		hasRow, err := statement.Step()
		if err != nil {
			*errRef = errors.WithStack(err)
			return
		}

		if !hasRow {
			break
		}

		id, _, err := statement.ColumnInt(0)
		if err != nil {
			*errRef = errors.WithStack(err)
			return
		}

		name, _, err := statement.ColumnRawString(1)
		if err != nil {
			*errRef = errors.WithStack(err)
			return
		}

//line pkg/server/json.qtpl:110
		if isFirstRow {
//line pkg/server/json.qtpl:111
			isFirstRow = false

//line pkg/server/json.qtpl:112
		} else {
//line pkg/server/json.qtpl:112
			qw422016.N().S(`,`)
//line pkg/server/json.qtpl:114
		}
//line pkg/server/json.qtpl:114
		qw422016.N().S(`{"id":`)
//line pkg/server/json.qtpl:116
		qw422016.N().D(id)
//line pkg/server/json.qtpl:116
		qw422016.N().S(`,"name":`)
//line pkg/server/json.qtpl:117
		qw422016.N().Q(string(name))
//line pkg/server/json.qtpl:117
		qw422016.N().S(`}`)
//line pkg/server/json.qtpl:119
	}
//line pkg/server/json.qtpl:119
	qw422016.N().S(`]`)
//line pkg/server/json.qtpl:121
}

//line pkg/server/json.qtpl:121
func writewriteMachineInfoList(qq422016 qtio422016.Writer, statement *sqlite3.Stmt, errRef *error) {
//line pkg/server/json.qtpl:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/server/json.qtpl:121
	streamwriteMachineInfoList(qw422016, statement, errRef)
//line pkg/server/json.qtpl:121
	qt422016.ReleaseWriter(qw422016)
//line pkg/server/json.qtpl:121
}

//line pkg/server/json.qtpl:121
func writeMachineInfoList(statement *sqlite3.Stmt, errRef *error) string {
//line pkg/server/json.qtpl:121
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/server/json.qtpl:121
	writewriteMachineInfoList(qb422016, statement, errRef)
//line pkg/server/json.qtpl:121
	qs422016 := string(qb422016.B)
//line pkg/server/json.qtpl:121
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/server/json.qtpl:121
	return qs422016
//line pkg/server/json.qtpl:121
}
