package com.saintrivers.rift.app.topic;

import java.util.List;

import com.saintrivers.rift.domain.Topic;
import com.saintrivers.rift.web.request.TopicRequest;

public interface TopicService {
    void create(TopicRequest topicRequest);

    List<Topic> getByCourseId(String courseId);

    void update(String topicId, TopicRequest topicRequest);

    void delete(String topicId);
}
