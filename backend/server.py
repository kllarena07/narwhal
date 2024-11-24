from fastapi import FastAPI, Form

app = FastAPI()


@app.post("/run")
def run(
        repo: str = Form(...),
        app_type: str = Form(...)
        ):
    print(repo, app_type)

    return "Hello World"


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)
