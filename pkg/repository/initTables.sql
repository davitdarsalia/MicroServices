CREATE TABLE "public.User" (
                               "userid" serial NOT NULL,
                               "personal_number" varchar(11) NOT NULL UNIQUE,
                               "phonenum" serial(15) NOT NULL,
                               "username" varchar(65) NOT NULL UNIQUE,
                               "email" varchar(100) NOT NULL UNIQUE,
                               "firstname" varchar(70) NOT NULL,
                               "lastname" varchar(70) NOT NULL,
                               "ip_address" varchar(30),
                               "password" varchar(128),
                               "salt" bytea(255) NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.balanceInfo" (
                                      "userid" int NOT NULL,
                                      "overallBalance" money NOT NULL,
                                      "debts" int NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.deposit" (
                                  "userid" int NOT NULL,
                                  "balance" int NOT NULL,
                                  "expiryDate" DATE NOT NULL,
                                  "addition_per_transaction" int NOT NULL,
                                  "currency_id" varchar(30) NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.rating" (
                                 "userid" int NOT NULL,
                                 "rating" int NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.receivedCashback" (
                                           "userid" int NOT NULL,
                                           "overallcashback" int NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.receivedBonus" (
                                        "userid" int NOT NULL,
                                        "overallBonus" int NOT NULL
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.countries" (
                                    "id" serial NOT NULL,
                                    "code" varchar(5) NOT NULL UNIQUE,
                                    "countryname" varchar(50) NOT NULL UNIQUE
) WITH (
      OIDS=FALSE
    );



CREATE TABLE "public.availableCurrencies" (
    "currency" varchar(5) NOT NULL UNIQUE
) WITH (
      OIDS=FALSE
    );




ALTER TABLE "balanceInfo" ADD CONSTRAINT "balanceInfo_fk0" FOREIGN KEY ("userid") REFERENCES "User"("userid");
ALTER TABLE "balanceInfo" ADD CONSTRAINT "balanceInfo_fk1" FOREIGN KEY ("debts") REFERENCES "User"("userid");

ALTER TABLE "deposit" ADD CONSTRAINT "deposit_fk0" FOREIGN KEY ("userid") REFERENCES "User"("userid");
ALTER TABLE "deposit" ADD CONSTRAINT "deposit_fk1" FOREIGN KEY ("currency_id") REFERENCES "availableCurrencies"("currency");

ALTER TABLE "rating" ADD CONSTRAINT "rating_fk0" FOREIGN KEY ("userid") REFERENCES "User"("userid");

ALTER TABLE "receivedCashback" ADD CONSTRAINT "receivedCashback_fk0" FOREIGN KEY ("userid") REFERENCES "User"("userid");

ALTER TABLE "receivedBonus" ADD CONSTRAINT "receivedBonus_fk0" FOREIGN KEY ("userid") REFERENCES "User"("userid");











