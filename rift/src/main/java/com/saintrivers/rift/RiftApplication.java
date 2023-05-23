package com.saintrivers.rift;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.mongodb.repository.config.EnableMongoRepositories;

@SpringBootApplication
@EnableMongoRepositories
public class RiftApplication {

	public static void main(String[] args) {
		SpringApplication.run(RiftApplication.class, args);
	}

}
