import time
import RPi.GPIO as GPIO

pin = 21

GPIO.setmode(GPIO.BCM)
GPIO.setup(pin, GPIO.IN)

while True:
	time.sleep(1)
	print GPIO.input(pin)
