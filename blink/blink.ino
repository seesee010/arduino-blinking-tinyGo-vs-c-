// blink an LED, c++ version
#define LED 13  // is this the right pin?

void setup() {
  pinMode(LED, OUTPUT);
}

void loop() {
  digitalWrite(LED, HIGH);
  delay(200);

  digitalWrite(LED, LOW);
  delay(200);
}
