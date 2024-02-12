// Code generated by hertz generator. DO NOT EDIT.

package line

import (
	line "GraduateThesis/biz/handler/line"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_line := _api.Group("/line", _lineMw()...)
			_line.POST("/bulk", append(_bulkMw(), line.Bulk)...)
			_line.POST("/change-dependence", append(_changedependenceMw(), line.ChangeDependence)...)
			_line.POST("/delete", append(_deleteMw(), line.Delete)...)
			_line.GET("/query", append(_queryMw(), line.Query)...)
			_line.POST("/topology_analyse", append(_topologyanalyseMw(), line.TopologyAnalyse)...)
		}
	}
}
