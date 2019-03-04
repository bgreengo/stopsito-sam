package com.dolphin.handler;

import com.amazonaws.handlers.RequestHandler2;
import com.dolphin.domain.ReservationRequest;
import com.dolphin.domain.ReservationResponse;
import com.dolphin.exception.ReservationMissingFieldsException;
import com.dolphin.service.PersistenceService;

public final class ReservationHandler extends RequestHandler2 {

    private static final PersistenceService persistenceService = new PersistenceService();

    public ReservationResponse handleRequest(ReservationRequest request) throws ReservationMissingFieldsException {
        if (null == request) {
            throw new IllegalStateException("Reservation Request cannot be null");
        }

        if (isRequestMissingRequiredFields(request)) {
            throw new ReservationMissingFieldsException();
        }
        return persistenceService.put(request);
    }

    private boolean isRequestMissingRequiredFields(ReservationRequest request) {
        return isRequestMissingNameOrEmailFields(request) || isRequestMissingPartyOrTimestampFields(request);
    }

    private boolean isRequestMissingNameOrEmailFields(ReservationRequest request) {
        return isStringBlank(request.getEmail()) || isStringBlank(request.getName());
    }

    private boolean isRequestMissingPartyOrTimestampFields(ReservationRequest request) {
        return request.getParty() < 1 || request.getTimestamp() < 1;
    }

    private boolean isStringBlank(String s) {
        return null == s || s.trim().length() == 0;
    }

}
