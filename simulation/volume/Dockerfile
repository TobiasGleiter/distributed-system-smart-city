FROM python:3.9-slim
WORKDIR /usr/app/src
COPY . .
RUN pip install --no-cache-dir -r requirements.txt
EXPOSE 4222
CMD ["python", "./main.py"]