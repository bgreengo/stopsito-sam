package domain;

public final class ReservationResponse {

    private final String one;

    public ReservationResponse(String one) {
        this.one = one;
    }

    public String getOne() {
        return one;
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("ReservationResponse{");
        sb.append("one='").append(one).append('\'');
        sb.append('}');
        return sb.toString();
    }
}
