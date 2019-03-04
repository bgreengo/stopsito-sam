package com.dolphin.domain;

import java.util.Objects;

// According to AWS: The `get` and `set` methods are required in order for the POJOs to work with AWS Lambda's built in JSON serializer.
// Which is horrible because you can't make these classes truly immutable :(
// https://docs.aws.amazon.com/lambda/latest/dg/java-handler-io-type-pojo.html

public final class ReservationRequest {

    private String name;
    private String phone;
    private String email;
    private long party;
    private long timestamp;
    private String message;

    public String getName() {
        return name;
    }

    public String getPhone() {
        return phone;
    }

    public String getEmail() {
        return email;
    }

    public long getParty() {
        return party;
    }

    public long getTimestamp() {
        return timestamp;
    }

    public String getMessage() {
        return message;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public void setParty(long party) {
        this.party = party;
    }

    public void setTimestamp(long timestamp) {
        this.timestamp = timestamp;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        ReservationRequest that = (ReservationRequest) o;
        return party == that.party &&
            Objects.equals(name, that.name) &&
            Objects.equals(phone, that.phone) &&
            Objects.equals(email, that.email) &&
            Objects.equals(timestamp, that.timestamp) &&
            Objects.equals(message, that.message);
    }

    @Override
    public int hashCode() {
        return Objects.hash(name, phone, email, party, timestamp, message);
    }

    @Override
    public String toString() {
        final StringBuffer sb = new StringBuffer("ReservationRequest{");
        sb.append("name='").append(name).append('\'');
        sb.append(", phone='").append(phone).append('\'');
        sb.append(", email='").append(email).append('\'');
        sb.append(", party=").append(party);
        sb.append(", timestamp=").append(timestamp);
        sb.append(", message='").append(message).append('\'');
        sb.append('}');
        return sb.toString();
    }
}
