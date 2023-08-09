FROM python:3.8-slim

WORKDIR /app
COPY sca .
RUN pip install -r ./requirements.txt
CMD ["python3", "main.py"]
