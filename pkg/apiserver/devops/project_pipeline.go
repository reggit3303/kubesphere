package devops

import (
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/errors"
	"kubesphere.io/kubesphere/pkg/models/devops"
	"net/http"
)

func CreateDevOpsProjectPipelineHandler(request *restful.Request, resp *restful.Response) {

	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	var pipeline *devops.ProjectPipeline
	err := request.ReadEntity(&pipeline)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusBadRequest, err.Error()), resp)
		return
	}
	err = devops.CheckProjectUserInRole(username, projectId, []string{devops.ProjectOwner, devops.ProjectMaintainer})
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	pipelineName, err := devops.CreateProjectPipeline(projectId, pipeline)

	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}

	resp.WriteAsJson(struct {
		Name string `json:"name"`
	}{Name: pipelineName})
	return
}

func DeleteDevOpsProjectPipelineHandler(request *restful.Request, resp *restful.Response) {
	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	pipelineId := request.PathParameter("pipelines")

	err := devops.CheckProjectUserInRole(username, projectId, []string{devops.ProjectOwner, devops.ProjectMaintainer})
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	pipelineName, err := devops.DeleteProjectPipeline(projectId, pipelineId)

	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}

	resp.WriteAsJson(struct {
		Name string `json:"name"`
	}{Name: pipelineName})
	return
}

func UpdateDevOpsProjectPipelineHandler(request *restful.Request, resp *restful.Response) {

	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	pipelineId := request.PathParameter("pipelines")
	var pipeline *devops.ProjectPipeline
	err := request.ReadEntity(&pipeline)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusBadRequest, err.Error()), resp)
		return
	}
	err = devops.CheckProjectUserInRole(username, projectId, []string{devops.ProjectOwner, devops.ProjectMaintainer})
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	pipelineName, err := devops.UpdateProjectPipeline(projectId, pipelineId, pipeline)

	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}

	resp.WriteAsJson(struct {
		Name string `json:"name"`
	}{Name: pipelineName})
	return
}

func GetDevOpsProjectPipelineHandler(request *restful.Request, resp *restful.Response) {

	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	pipelineId := request.PathParameter("pipelines")

	err := devops.CheckProjectUserInRole(username, projectId, []string{devops.ProjectOwner, devops.ProjectMaintainer})
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	pipeline, err := devops.GetProjectPipeline(projectId, pipelineId)

	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}

	resp.WriteAsJson(pipeline)
	return
}

func GetPipelineSonarStatusHandler(request *restful.Request, resp *restful.Response) {
	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	pipelineId := request.PathParameter("pipelines")
	err := devops.CheckProjectUserInRole(username, projectId, devops.AllRoleSlice)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	sonarStatus, err := devops.GetPipelineSonar(projectId, pipelineId)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}
	resp.WriteAsJson(sonarStatus)
}

func GetMultiBranchesPipelineSonarStatusHandler(request *restful.Request, resp *restful.Response) {
	projectId := request.PathParameter("devops")
	username := request.HeaderParameter(constants.UserNameHeader)
	pipelineId := request.PathParameter("pipelines")
	branchId := request.PathParameter("branches")
	err := devops.CheckProjectUserInRole(username, projectId, devops.AllRoleSlice)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(restful.NewError(http.StatusForbidden, err.Error()), resp)
		return
	}
	sonarStatus, err := devops.GetMultiBranchPipelineSonar(projectId, pipelineId, branchId)
	if err != nil {
		glog.Errorf("%+v", err)
		errors.ParseSvcErr(err, resp)
		return
	}
	resp.WriteAsJson(sonarStatus)
}
