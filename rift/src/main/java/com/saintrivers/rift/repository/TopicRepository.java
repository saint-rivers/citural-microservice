package com.saintrivers.rift.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.saintrivers.rift.domain.Topic;

public interface TopicRepository extends MongoRepository<Topic, String> {
    
}
