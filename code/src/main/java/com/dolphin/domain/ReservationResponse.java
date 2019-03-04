package com.dolphin.domain;

import java.util.Objects;

public final class ReservationResponse {

    private final String id;
    private final long timestamp;

    public ReservationResponse() {
        this("", 0);
    }

    public ReservationResponse(String id, long timestamp) {
        this.id = id;
        this.timestamp = timestamp;
    }

    public String getId() {
        return id;
    }

    public long getTimestamp() {
        return timestamp;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        ReservationResponse that = (ReservationResponse) o;
        return Objects.equals(id, that.id) &&
            Objects.equals(timestamp, that.timestamp);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, timestamp);
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("ReservationResponse{");
        sb.append("id='").append(id).append('\'');
        sb.append(", timestamp=").append(timestamp);
        sb.append('}');
        return sb.toString();
    }
}
