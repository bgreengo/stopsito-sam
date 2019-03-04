package com.dolphin.handler;

import com.amazonaws.handlers.RequestHandler2;
import com.dolphin.domain.ReservationRequest;
import com.dolphin.domain.ReservationResponse;
import com.dolphin.exception.ReservationMissingFieldsException;
import java.util.UUID;

public final class ReservationHandler extends RequestHandler2 {

    public ReservationResponse handleRequest(ReservationRequest request) throws ReservationMissingFieldsException {
        if (null == request) {
            throw new IllegalStateException("Reservation Request cannot be null");
        }

        if (isRequestMissingNameOrEmailFields(request)) {
            throw new ReservationMissingFieldsException();
        }

        

        return new ReservationResponse(UUID.randomUUID().toString(), request.getTimestamp());
    }

    public boolean isRequestMissingNameOrEmailFields(ReservationRequest request) {
        return isStringBlank(request.getEmail()) || isStringBlank(request.getName());
    }

    private boolean isStringBlank(String email) {
        return null == email || email.trim().length() == 0;
    }

}
