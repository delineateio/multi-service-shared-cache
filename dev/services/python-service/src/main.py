from fastapi import FastAPI, Response,status
import socket
import cache

app = FastAPI()

@app.get("/healthz", status_code=200, response_class=Response)
async def healthz(response: Response):
  addHeaders(response)


@app.post("/increment", status_code=202, response_class=Response)
async def increment(response: Response):
  addHeaders(response)
  cache.write()


@app.get("/count", status_code=200)
async def count(response: Response):
  addHeaders(response)
  count = cache.read()
  return {
    "count": count
  }

def addHeaders(response):
  host = socket.gethostname()
  response.headers["delineateio-hostname"] = host
  response.headers["delineateio-lang"] = "python"
