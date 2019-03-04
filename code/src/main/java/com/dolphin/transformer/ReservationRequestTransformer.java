package com.dolphin.transformer;

import com.amazonaws.services.dynamodbv2.document.Item;
import com.dolphin.domain.ReservationRequest;
import java.util.Optional;
import java.util.UUID;

public final class ReservationRequestTransformer {

    public Item transform(ReservationRequest request) {

        final Item item = new Item()
            .withPrimaryKey("ID", UUID.randomUUID().toString())
            .withString("name", request.getName())
            .withString("email", request.getEmail())
            .withLong("timestamp", request.getTimestamp())
            .withLong("party", request.getParty());

        fieldIsPresent(request.getPhone()).ifPresent(phone -> item.withString("phone", phone));
        fieldIsPresent(request.getMessage()).ifPresent(message -> item.withString("message", message));

        return item;
    }

    private Optional<String> fieldIsPresent(String s) {
        return (null != s && s.trim().length() > 0) ? Optional.of(s) : Optional.empty();
    }

}
