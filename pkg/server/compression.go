package server

import (
	"bytes"
	dataquery "github.com/JetBrains/ij-perf-report-aggregator/pkg/data-query"
	"github.com/JetBrains/ij-perf-report-aggregator/pkg/util"
	"github.com/develar/errors"
	"github.com/klauspost/compress/zstd"
	"github.com/valyala/bytebufferpool"
)

func decompressData(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	reader, err := zstd.NewReader(bytes.NewReader(input), zstd.WithDecoderDicts(dataquery.ZstdDictionary))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return CopyBuffer(buf), nil
}

func (rcm *ResponseCacheManager) compressData(value []byte) ([]byte, error) {
	buffer := bytebufferpool.Get()
	defer bytebufferpool.Put(buffer)
	writer, err := zstd.NewWriter(buffer, zstd.WithEncoderLevel(zstd.SpeedFastest), zstd.WithEncoderDict(dataquery.ZstdDictionary))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = writer.Write(value)
	if err != nil {
		util.Close(writer, rcm.logger)
		return nil, err
	}

	util.Close(writer, rcm.logger)
	return CopyBuffer(buffer), nil
}
