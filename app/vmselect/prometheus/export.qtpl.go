// Code generated by qtc from "export.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/export.qtpl:1
package prometheus

//line app/vmselect/prometheus/export.qtpl:1
import (
	"bytes"
	"math"
	"strings"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
	"github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/export.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/export.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/export.qtpl:13
func StreamExportCSVLine(qw422016 *qt422016.Writer, xb *exportBlock, fieldNames []string) {
//line app/vmselect/prometheus/export.qtpl:14
	if len(xb.timestamps) == 0 || len(fieldNames) == 0 {
//line app/vmselect/prometheus/export.qtpl:14
		return
//line app/vmselect/prometheus/export.qtpl:14
	}
//line app/vmselect/prometheus/export.qtpl:15
	for i, timestamp := range xb.timestamps {
//line app/vmselect/prometheus/export.qtpl:16
		value := xb.values[i]

//line app/vmselect/prometheus/export.qtpl:17
		streamexportCSVField(qw422016, xb.mn, fieldNames[0], timestamp, value)
//line app/vmselect/prometheus/export.qtpl:18
		for _, fieldName := range fieldNames[1:] {
//line app/vmselect/prometheus/export.qtpl:18
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:20
			streamexportCSVField(qw422016, xb.mn, fieldName, timestamp, value)
//line app/vmselect/prometheus/export.qtpl:21
		}
//line app/vmselect/prometheus/export.qtpl:22
		qw422016.N().S(`
`)
//line app/vmselect/prometheus/export.qtpl:23
	}
//line app/vmselect/prometheus/export.qtpl:24
}

//line app/vmselect/prometheus/export.qtpl:24
func WriteExportCSVLine(qq422016 qtio422016.Writer, xb *exportBlock, fieldNames []string) {
//line app/vmselect/prometheus/export.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:24
	StreamExportCSVLine(qw422016, xb, fieldNames)
//line app/vmselect/prometheus/export.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:24
}

//line app/vmselect/prometheus/export.qtpl:24
func ExportCSVLine(xb *exportBlock, fieldNames []string) string {
//line app/vmselect/prometheus/export.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:24
	WriteExportCSVLine(qb422016, xb, fieldNames)
//line app/vmselect/prometheus/export.qtpl:24
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:24
	return qs422016
//line app/vmselect/prometheus/export.qtpl:24
}

//line app/vmselect/prometheus/export.qtpl:26
func streamexportCSVField(qw422016 *qt422016.Writer, mn *storage.MetricName, fieldName string, timestamp int64, value float64) {
//line app/vmselect/prometheus/export.qtpl:27
	if fieldName == "__value__" {
//line app/vmselect/prometheus/export.qtpl:28
		qw422016.N().F(value)
//line app/vmselect/prometheus/export.qtpl:29
		return
//line app/vmselect/prometheus/export.qtpl:30
	}
//line app/vmselect/prometheus/export.qtpl:31
	if fieldName == "__timestamp__" {
//line app/vmselect/prometheus/export.qtpl:32
		qw422016.N().DL(timestamp)
//line app/vmselect/prometheus/export.qtpl:33
		return
//line app/vmselect/prometheus/export.qtpl:34
	}
//line app/vmselect/prometheus/export.qtpl:35
	if strings.HasPrefix(fieldName, "__timestamp__:") {
//line app/vmselect/prometheus/export.qtpl:36
		timeFormat := fieldName[len("__timestamp__:"):]

//line app/vmselect/prometheus/export.qtpl:37
		switch timeFormat {
//line app/vmselect/prometheus/export.qtpl:38
		case "unix_s":
//line app/vmselect/prometheus/export.qtpl:39
			qw422016.N().DL(timestamp / 1000)
//line app/vmselect/prometheus/export.qtpl:40
		case "unix_ms":
//line app/vmselect/prometheus/export.qtpl:41
			qw422016.N().DL(timestamp)
//line app/vmselect/prometheus/export.qtpl:42
		case "unix_ns":
//line app/vmselect/prometheus/export.qtpl:43
			qw422016.N().DL(timestamp * 1e6)
//line app/vmselect/prometheus/export.qtpl:44
		case "rfc3339":
//line app/vmselect/prometheus/export.qtpl:46
			bb := quicktemplate.AcquireByteBuffer()
			bb.B = time.Unix(timestamp/1000, (timestamp%1000)*1e6).AppendFormat(bb.B[:0], time.RFC3339)

//line app/vmselect/prometheus/export.qtpl:49
			qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:51
			quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:53
		default:
//line app/vmselect/prometheus/export.qtpl:54
			if strings.HasPrefix(timeFormat, "custom:") {
//line app/vmselect/prometheus/export.qtpl:56
				layout := timeFormat[len("custom:"):]
				bb := quicktemplate.AcquireByteBuffer()
				bb.B = time.Unix(timestamp/1000, (timestamp%1000)*1e6).AppendFormat(bb.B[:0], layout)

//line app/vmselect/prometheus/export.qtpl:60
				if bytes.ContainsAny(bb.B, `"`+",\n") {
//line app/vmselect/prometheus/export.qtpl:61
					qw422016.E().QZ(bb.B)
//line app/vmselect/prometheus/export.qtpl:62
				} else {
//line app/vmselect/prometheus/export.qtpl:63
					qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:64
				}
//line app/vmselect/prometheus/export.qtpl:66
				quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:68
			} else {
//line app/vmselect/prometheus/export.qtpl:68
				qw422016.N().S(`Unsupported timeFormat=`)
//line app/vmselect/prometheus/export.qtpl:69
				qw422016.N().S(timeFormat)
//line app/vmselect/prometheus/export.qtpl:70
			}
//line app/vmselect/prometheus/export.qtpl:71
		}
//line app/vmselect/prometheus/export.qtpl:72
		return
//line app/vmselect/prometheus/export.qtpl:73
	}
//line app/vmselect/prometheus/export.qtpl:74
	v := mn.GetTagValue(fieldName)

//line app/vmselect/prometheus/export.qtpl:75
	if bytes.ContainsAny(v, `"`+",\n") {
//line app/vmselect/prometheus/export.qtpl:76
		qw422016.N().QZ(v)
//line app/vmselect/prometheus/export.qtpl:77
	} else {
//line app/vmselect/prometheus/export.qtpl:78
		qw422016.N().Z(v)
//line app/vmselect/prometheus/export.qtpl:79
	}
//line app/vmselect/prometheus/export.qtpl:80
}

//line app/vmselect/prometheus/export.qtpl:80
func writeexportCSVField(qq422016 qtio422016.Writer, mn *storage.MetricName, fieldName string, timestamp int64, value float64) {
//line app/vmselect/prometheus/export.qtpl:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:80
	streamexportCSVField(qw422016, mn, fieldName, timestamp, value)
//line app/vmselect/prometheus/export.qtpl:80
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:80
}

//line app/vmselect/prometheus/export.qtpl:80
func exportCSVField(mn *storage.MetricName, fieldName string, timestamp int64, value float64) string {
//line app/vmselect/prometheus/export.qtpl:80
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:80
	writeexportCSVField(qb422016, mn, fieldName, timestamp, value)
//line app/vmselect/prometheus/export.qtpl:80
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:80
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:80
	return qs422016
//line app/vmselect/prometheus/export.qtpl:80
}

//line app/vmselect/prometheus/export.qtpl:82
func StreamExportPrometheusLine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:83
	if len(xb.timestamps) == 0 {
//line app/vmselect/prometheus/export.qtpl:83
		return
//line app/vmselect/prometheus/export.qtpl:83
	}
//line app/vmselect/prometheus/export.qtpl:84
	bb := quicktemplate.AcquireByteBuffer()

//line app/vmselect/prometheus/export.qtpl:85
	writeprometheusMetricName(bb, xb.mn)

//line app/vmselect/prometheus/export.qtpl:86
	for i, ts := range xb.timestamps {
//line app/vmselect/prometheus/export.qtpl:87
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:87
		qw422016.N().S(` `)
//line app/vmselect/prometheus/export.qtpl:88
		qw422016.N().F(xb.values[i])
//line app/vmselect/prometheus/export.qtpl:88
		qw422016.N().S(` `)
//line app/vmselect/prometheus/export.qtpl:89
		qw422016.N().DL(ts)
//line app/vmselect/prometheus/export.qtpl:89
		qw422016.N().S(`
`)
//line app/vmselect/prometheus/export.qtpl:90
	}
//line app/vmselect/prometheus/export.qtpl:91
	quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:92
}

//line app/vmselect/prometheus/export.qtpl:92
func WriteExportPrometheusLine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:92
	StreamExportPrometheusLine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:92
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:92
}

//line app/vmselect/prometheus/export.qtpl:92
func ExportPrometheusLine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:92
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:92
	WriteExportPrometheusLine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:92
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:92
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:92
	return qs422016
//line app/vmselect/prometheus/export.qtpl:92
}

//line app/vmselect/prometheus/export.qtpl:94
func StreamExportJSONLine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:95
	if len(xb.timestamps) == 0 {
//line app/vmselect/prometheus/export.qtpl:95
		return
//line app/vmselect/prometheus/export.qtpl:95
	}
//line app/vmselect/prometheus/export.qtpl:95
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/export.qtpl:97
	streammetricNameObject(qw422016, xb.mn)
//line app/vmselect/prometheus/export.qtpl:97
	qw422016.N().S(`,"values":[`)
//line app/vmselect/prometheus/export.qtpl:99
	if len(xb.values) > 0 {
//line app/vmselect/prometheus/export.qtpl:100
		values := xb.values

//line app/vmselect/prometheus/export.qtpl:101
		qw422016.N().F(values[0])
//line app/vmselect/prometheus/export.qtpl:102
		values = values[1:]

//line app/vmselect/prometheus/export.qtpl:103
		for _, v := range values {
//line app/vmselect/prometheus/export.qtpl:103
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:104
			if math.IsNaN(v) {
//line app/vmselect/prometheus/export.qtpl:104
				qw422016.N().S(`null`)
//line app/vmselect/prometheus/export.qtpl:104
			} else {
//line app/vmselect/prometheus/export.qtpl:104
				qw422016.N().F(v)
//line app/vmselect/prometheus/export.qtpl:104
			}
//line app/vmselect/prometheus/export.qtpl:105
		}
//line app/vmselect/prometheus/export.qtpl:106
	}
//line app/vmselect/prometheus/export.qtpl:106
	qw422016.N().S(`],"timestamps":[`)
//line app/vmselect/prometheus/export.qtpl:109
	if len(xb.timestamps) > 0 {
//line app/vmselect/prometheus/export.qtpl:110
		timestamps := xb.timestamps

//line app/vmselect/prometheus/export.qtpl:111
		qw422016.N().DL(timestamps[0])
//line app/vmselect/prometheus/export.qtpl:112
		timestamps = timestamps[1:]

//line app/vmselect/prometheus/export.qtpl:113
		for _, ts := range timestamps {
//line app/vmselect/prometheus/export.qtpl:113
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:114
			qw422016.N().DL(ts)
//line app/vmselect/prometheus/export.qtpl:115
		}
//line app/vmselect/prometheus/export.qtpl:116
	}
//line app/vmselect/prometheus/export.qtpl:116
	qw422016.N().S(`]}`)
//line app/vmselect/prometheus/export.qtpl:118
	qw422016.N().S(`
`)
//line app/vmselect/prometheus/export.qtpl:119
}

//line app/vmselect/prometheus/export.qtpl:119
func WriteExportJSONLine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:119
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:119
	StreamExportJSONLine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:119
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:119
}

//line app/vmselect/prometheus/export.qtpl:119
func ExportJSONLine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:119
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:119
	WriteExportJSONLine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:119
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:119
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:119
	return qs422016
//line app/vmselect/prometheus/export.qtpl:119
}

//line app/vmselect/prometheus/export.qtpl:121
func StreamExportPromAPILine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:121
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/export.qtpl:123
	streammetricNameObject(qw422016, xb.mn)
//line app/vmselect/prometheus/export.qtpl:123
	qw422016.N().S(`,"values":`)
//line app/vmselect/prometheus/export.qtpl:124
	streamvaluesWithTimestamps(qw422016, xb.values, xb.timestamps)
//line app/vmselect/prometheus/export.qtpl:124
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/export.qtpl:126
}

//line app/vmselect/prometheus/export.qtpl:126
func WriteExportPromAPILine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:126
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:126
	StreamExportPromAPILine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:126
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:126
}

//line app/vmselect/prometheus/export.qtpl:126
func ExportPromAPILine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:126
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:126
	WriteExportPromAPILine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:126
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:126
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:126
	return qs422016
//line app/vmselect/prometheus/export.qtpl:126
}

//line app/vmselect/prometheus/export.qtpl:128
func StreamExportPromAPIResponse(qw422016 *qt422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:128
	qw422016.N().S(`{"status":"success","data":{"resultType":"matrix","result":[`)
//line app/vmselect/prometheus/export.qtpl:134
	bb, ok := <-resultsCh

//line app/vmselect/prometheus/export.qtpl:135
	if ok {
//line app/vmselect/prometheus/export.qtpl:136
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:137
		quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:138
		for bb := range resultsCh {
//line app/vmselect/prometheus/export.qtpl:138
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:139
			qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:140
			quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:141
		}
//line app/vmselect/prometheus/export.qtpl:142
	}
//line app/vmselect/prometheus/export.qtpl:142
	qw422016.N().S(`]}}`)
//line app/vmselect/prometheus/export.qtpl:146
}

//line app/vmselect/prometheus/export.qtpl:146
func WriteExportPromAPIResponse(qq422016 qtio422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:146
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:146
	StreamExportPromAPIResponse(qw422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:146
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:146
}

//line app/vmselect/prometheus/export.qtpl:146
func ExportPromAPIResponse(resultsCh <-chan *quicktemplate.ByteBuffer) string {
//line app/vmselect/prometheus/export.qtpl:146
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:146
	WriteExportPromAPIResponse(qb422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:146
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:146
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:146
	return qs422016
//line app/vmselect/prometheus/export.qtpl:146
}

//line app/vmselect/prometheus/export.qtpl:148
func StreamExportStdResponse(qw422016 *qt422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:149
	for bb := range resultsCh {
//line app/vmselect/prometheus/export.qtpl:150
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:151
		quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:152
	}
//line app/vmselect/prometheus/export.qtpl:153
}

//line app/vmselect/prometheus/export.qtpl:153
func WriteExportStdResponse(qq422016 qtio422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:153
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:153
	StreamExportStdResponse(qw422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:153
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:153
}

//line app/vmselect/prometheus/export.qtpl:153
func ExportStdResponse(resultsCh <-chan *quicktemplate.ByteBuffer) string {
//line app/vmselect/prometheus/export.qtpl:153
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:153
	WriteExportStdResponse(qb422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:153
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:153
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:153
	return qs422016
//line app/vmselect/prometheus/export.qtpl:153
}

//line app/vmselect/prometheus/export.qtpl:155
func streamprometheusMetricName(qw422016 *qt422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/export.qtpl:156
	qw422016.N().Z(mn.MetricGroup)
//line app/vmselect/prometheus/export.qtpl:157
	if len(mn.Tags) > 0 {
//line app/vmselect/prometheus/export.qtpl:157
		qw422016.N().S(`{`)
//line app/vmselect/prometheus/export.qtpl:159
		tags := mn.Tags

//line app/vmselect/prometheus/export.qtpl:160
		qw422016.N().Z(tags[0].Key)
//line app/vmselect/prometheus/export.qtpl:160
		qw422016.N().S(`=`)
//line app/vmselect/prometheus/export.qtpl:160
		qw422016.N().QZ(tags[0].Value)
//line app/vmselect/prometheus/export.qtpl:161
		tags = tags[1:]

//line app/vmselect/prometheus/export.qtpl:162
		for i := range tags {
//line app/vmselect/prometheus/export.qtpl:163
			tag := &tags[i]

//line app/vmselect/prometheus/export.qtpl:163
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:164
			qw422016.N().Z(tag.Key)
//line app/vmselect/prometheus/export.qtpl:164
			qw422016.N().S(`=`)
//line app/vmselect/prometheus/export.qtpl:164
			qw422016.N().QZ(tag.Value)
//line app/vmselect/prometheus/export.qtpl:165
		}
//line app/vmselect/prometheus/export.qtpl:165
		qw422016.N().S(`}`)
//line app/vmselect/prometheus/export.qtpl:167
	}
//line app/vmselect/prometheus/export.qtpl:168
}

//line app/vmselect/prometheus/export.qtpl:168
func writeprometheusMetricName(qq422016 qtio422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/export.qtpl:168
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:168
	streamprometheusMetricName(qw422016, mn)
//line app/vmselect/prometheus/export.qtpl:168
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:168
}

//line app/vmselect/prometheus/export.qtpl:168
func prometheusMetricName(mn *storage.MetricName) string {
//line app/vmselect/prometheus/export.qtpl:168
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:168
	writeprometheusMetricName(qb422016, mn)
//line app/vmselect/prometheus/export.qtpl:168
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:168
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:168
	return qs422016
//line app/vmselect/prometheus/export.qtpl:168
}
