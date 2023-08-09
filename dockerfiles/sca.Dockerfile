FROM python:3.8-slim

WORKDIR /app
COPY parse .
RUN pip install -f ./requirements.txt
CMD ["python3", "main.py"]
