package com.saintrivers.rift.app.topic;

import java.time.LocalDateTime;

import org.springframework.stereotype.Component;

import com.saintrivers.rift.domain.Topic;
import com.saintrivers.rift.web.request.TopicRequest;

@Component
public class TopicMapper {

    public Topic toDomain(TopicRequest request) {
        Topic topic = new Topic();
        topic.setCourseId(request.courseId());
        topic.setSections(request.sections());
        topic.setIsDeleted(false);
        return topic;
    }

    private Topic cleanDomain(Topic fetchedTopic) {
        fetchedTopic.setId(null);
        fetchedTopic.setChanges(null);
        fetchedTopic.setCourseId(null);
        fetchedTopic.setDateCreated(null);
        fetchedTopic.setIsDeleted(null);
        return fetchedTopic;
    }

    public Topic prepareUpdate(Topic fetchedTopic, Topic request) {
        var oldChanges = fetchedTopic.getChanges();
        fetchedTopic = this.cleanDomain(fetchedTopic);

        request.setChanges(oldChanges);
        request.logChanges(fetchedTopic);
        request.setLastUpdated(LocalDateTime.now());
        return request;
    }

}
