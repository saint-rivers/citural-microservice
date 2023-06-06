package com.saintrivers.commandtower.module.mock;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@RestController
@RequiredArgsConstructor
@RequestMapping("/mock")
@Slf4j 
public class MockController {

    @RequestMapping(path = "/project/{projectId}/**")
    public ResponseEntity<?> executeMock(@PathVariable int projectId, HttpServletRequest r){
        String fullUrl = r.getRequestURL().toString();
        log.info(fullUrl);

        return ResponseEntity.ok().build();
    }
}
