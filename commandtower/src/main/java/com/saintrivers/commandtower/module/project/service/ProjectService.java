package com.saintrivers.commandtower.module.project.service;

import java.util.List;

import com.saintrivers.commandtower.module.project.web.request.EndpointRequest;
import com.saintrivers.commandtower.module.project.web.request.ProjectRequest;
import com.saintrivers.commandtower.module.project.web.response.EndpointResponse;
import com.saintrivers.commandtower.module.project.web.response.ProjectResponse;
import com.saintrivers.commandtower.shared.response.PagedResponse;

public interface ProjectService {
    
    void create(ProjectRequest r);
    PagedResponse<List<ProjectResponse>> getPagedData(int page, int size);
    void addEndpoint(EndpointRequest endpointRequest);
    List<EndpointResponse>  getEndpoints(long projectId);
}
