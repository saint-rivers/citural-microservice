package com.saintrivers.rift.app.topic;

import java.util.List;
import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import com.saintrivers.rift.domain.Topic;

public interface TopicRepository extends MongoRepository<Topic, String> {

    @Query("{'courseId': ?0}")
    List<Topic> selectWhereCourseIdIs(String courseId);

    @Query("{'topicId': ?0}")
    Optional<Topic> findByTopicId(String topicId);
}
