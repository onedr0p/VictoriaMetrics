// Code generated by qtc from "query_range_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/query_range_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/query_range_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/promql"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/querytracer"
)

// QueryRangeResponse generates response for /api/v1/query_range.See https://prometheus.io/docs/prometheus/latest/querying/api/#range-queries

//line app/vmselect/prometheus/query_range_response.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/query_range_response.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/query_range_response.qtpl:10
func StreamQueryRangeResponse(qw422016 *qt422016.Writer, isPartial bool, rs []netstorage.Result, qt *querytracer.Tracer, qtDone func(), qs *promql.QueryStats) {
//line app/vmselect/prometheus/query_range_response.qtpl:10
	qw422016.N().S(`{`)
//line app/vmselect/prometheus/query_range_response.qtpl:13
	seriesCount := len(rs)
	pointsCount := 0

//line app/vmselect/prometheus/query_range_response.qtpl:15
	qw422016.N().S(`"status":"success","isPartial":`)
//line app/vmselect/prometheus/query_range_response.qtpl:17
	if isPartial {
//line app/vmselect/prometheus/query_range_response.qtpl:17
		qw422016.N().S(`true`)
//line app/vmselect/prometheus/query_range_response.qtpl:17
	} else {
//line app/vmselect/prometheus/query_range_response.qtpl:17
		qw422016.N().S(`false`)
//line app/vmselect/prometheus/query_range_response.qtpl:17
	}
//line app/vmselect/prometheus/query_range_response.qtpl:17
	qw422016.N().S(`,"data":{"resultType":"matrix","result":[`)
//line app/vmselect/prometheus/query_range_response.qtpl:21
	if len(rs) > 0 {
//line app/vmselect/prometheus/query_range_response.qtpl:22
		streamqueryRangeLine(qw422016, &rs[0])
//line app/vmselect/prometheus/query_range_response.qtpl:23
		pointsCount += len(rs[0].Values)

//line app/vmselect/prometheus/query_range_response.qtpl:24
		rs = rs[1:]

//line app/vmselect/prometheus/query_range_response.qtpl:25
		for i := range rs {
//line app/vmselect/prometheus/query_range_response.qtpl:25
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/query_range_response.qtpl:26
			streamqueryRangeLine(qw422016, &rs[i])
//line app/vmselect/prometheus/query_range_response.qtpl:27
			pointsCount += len(rs[i].Values)

//line app/vmselect/prometheus/query_range_response.qtpl:28
		}
//line app/vmselect/prometheus/query_range_response.qtpl:29
	}
//line app/vmselect/prometheus/query_range_response.qtpl:29
	qw422016.N().S(`]},"stats":{`)
//line app/vmselect/prometheus/query_range_response.qtpl:34
	// seriesFetched is string instead of int because of historical reasons.
	// It cannot be converted to int without breaking backwards compatibility at vmalert :(

//line app/vmselect/prometheus/query_range_response.qtpl:36
	qw422016.N().S(`"seriesFetched": "`)
//line app/vmselect/prometheus/query_range_response.qtpl:37
	qw422016.N().DL(qs.SeriesFetched)
//line app/vmselect/prometheus/query_range_response.qtpl:37
	qw422016.N().S(`","executionTimeMsec":`)
//line app/vmselect/prometheus/query_range_response.qtpl:38
	qw422016.N().DL(qs.ExecutionTimeMsec)
//line app/vmselect/prometheus/query_range_response.qtpl:38
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:41
	qt.Printf("generate /api/v1/query_range response for series=%d, points=%d", seriesCount, pointsCount)
	qtDone()

//line app/vmselect/prometheus/query_range_response.qtpl:44
	streamdumpQueryTrace(qw422016, qt)
//line app/vmselect/prometheus/query_range_response.qtpl:44
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:46
}

//line app/vmselect/prometheus/query_range_response.qtpl:46
func WriteQueryRangeResponse(qq422016 qtio422016.Writer, isPartial bool, rs []netstorage.Result, qt *querytracer.Tracer, qtDone func(), qs *promql.QueryStats) {
//line app/vmselect/prometheus/query_range_response.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:46
	StreamQueryRangeResponse(qw422016, isPartial, rs, qt, qtDone, qs)
//line app/vmselect/prometheus/query_range_response.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:46
}

//line app/vmselect/prometheus/query_range_response.qtpl:46
func QueryRangeResponse(isPartial bool, rs []netstorage.Result, qt *querytracer.Tracer, qtDone func(), qs *promql.QueryStats) string {
//line app/vmselect/prometheus/query_range_response.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:46
	WriteQueryRangeResponse(qb422016, isPartial, rs, qt, qtDone, qs)
//line app/vmselect/prometheus/query_range_response.qtpl:46
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:46
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:46
}

//line app/vmselect/prometheus/query_range_response.qtpl:48
func streamqueryRangeLine(qw422016 *qt422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:48
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/query_range_response.qtpl:50
	streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/prometheus/query_range_response.qtpl:50
	qw422016.N().S(`,"values":`)
//line app/vmselect/prometheus/query_range_response.qtpl:51
	streamvaluesWithTimestamps(qw422016, r.Values, r.Timestamps)
//line app/vmselect/prometheus/query_range_response.qtpl:51
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_range_response.qtpl:53
}

//line app/vmselect/prometheus/query_range_response.qtpl:53
func writequeryRangeLine(qq422016 qtio422016.Writer, r *netstorage.Result) {
//line app/vmselect/prometheus/query_range_response.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_range_response.qtpl:53
	streamqueryRangeLine(qw422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_range_response.qtpl:53
}

//line app/vmselect/prometheus/query_range_response.qtpl:53
func queryRangeLine(r *netstorage.Result) string {
//line app/vmselect/prometheus/query_range_response.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_range_response.qtpl:53
	writequeryRangeLine(qb422016, r)
//line app/vmselect/prometheus/query_range_response.qtpl:53
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_range_response.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_range_response.qtpl:53
	return qs422016
//line app/vmselect/prometheus/query_range_response.qtpl:53
}
