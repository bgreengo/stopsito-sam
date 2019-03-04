package com.dolphin.exception;

public class ReservationMissingFieldsException extends Exception {

    public ReservationMissingFieldsException() {
        super("ReservationRequest is missing fields. Mandatory fields: Name, email");
    }
}
