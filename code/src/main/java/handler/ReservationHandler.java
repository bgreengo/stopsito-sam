package handler;

import com.amazonaws.handlers.RequestHandler2;
import domain.ReservationRequest;
import domain.ReservationResponse;

public final class ReservationHandler extends RequestHandler2 {

    public ReservationResponse handleRequest(ReservationRequest reservationRequest) {
        return new ReservationResponse("egg");
    }

}
