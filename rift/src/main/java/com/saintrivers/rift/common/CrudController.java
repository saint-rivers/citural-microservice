package com.saintrivers.rift.common;

import java.util.List;

import org.springframework.web.bind.annotation.RequestParam;

import com.saintrivers.rift.common.response.BaseResponse;

public interface CrudController<T, R> {
    BaseResponse<?> create(R request);

    List<R> find(@RequestParam Integer page, @RequestParam Integer size);

    BaseResponse<?> update(String id, R request);
}
