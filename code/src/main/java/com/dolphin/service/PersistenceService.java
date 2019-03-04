package com.dolphin.service;

import com.amazonaws.services.dynamodbv2.AmazonDynamoDB;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDBClientBuilder;
import com.amazonaws.services.dynamodbv2.document.DynamoDB;
import com.amazonaws.services.dynamodbv2.document.Item;
import com.amazonaws.services.dynamodbv2.document.Table;
import com.dolphin.domain.ReservationRequest;
import com.dolphin.domain.ReservationResponse;
import com.dolphin.transformer.ReservationRequestTransformer;

public final class PersistenceService {

    private static final AmazonDynamoDB client = AmazonDynamoDBClientBuilder.standard().build();
    private static final DynamoDB dynamoDB = new DynamoDB(client);
    private static final Table reservationsTable = dynamoDB.getTable("Reservations");

    private static final ReservationRequestTransformer transformer = new ReservationRequestTransformer();

    public ReservationResponse put(ReservationRequest request) {
        final Item item = transformer.transform(request);
        reservationsTable.putItem(item);
        return new ReservationResponse(item.getString("ID"));
    }

}
