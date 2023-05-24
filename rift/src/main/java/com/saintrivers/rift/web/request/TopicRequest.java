package com.saintrivers.rift.web.request;

import java.util.List;

import com.saintrivers.rift.domain.model.Section;

import lombok.Builder;

@Builder
public record TopicRequest(
        List<Section> sections,
        String courseId) {
}
