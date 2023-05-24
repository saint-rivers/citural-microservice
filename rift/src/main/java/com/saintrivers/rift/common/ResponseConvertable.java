package com.saintrivers.rift.common;

public interface ResponseConvertable<D, R> {
    R toResponse(D domain);
}
