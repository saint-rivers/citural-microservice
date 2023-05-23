package com.saintrivers.rift.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.saintrivers.rift.domain.Course;

public interface CourseRepository extends MongoRepository<Course, String> {

}
