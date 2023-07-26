# AImate 

## Introduction
Welcome to AImate, a powerful tool that utilizes the OP Stack (OpenAI + Pinecone Vector Database) to enable users to upload their own custom knowledgebase files and ask questions about their contents.

**Website**: 



With AImate, you can quickly set up your own Golang server along with a user-friendly React frontend, allowing users to ask OpenAI questions about the specific knowledge base provided. The primary focus is on human-readable content like books, letters, and other documents, making it a practical and valuable tool for knowledge extraction and question-answering. Users can upload an entire library's worth of books and documents and receive pointed answers along with the name of the file and specific section within the file that the answer is based on!

## Features

With AImate, you can:

- Upload a variety of popular document types via a simple React frontend to create a custom knowledge base.
- Retrieve accurate and relevant answers based on the content of your uploaded documents.
- View the filenames and specific context snippets that inform the answer.
- Explore the power of the OP Stack (OpenAI + Pinecone Vector Database) in a user-friendly interface.
- Load entire libraries' worth of books into AImate.

## Manual Dependencies

Before getting started, ensure that you have the following dependencies installed:

- Node v19
- Go v1.18.9 (darwin/arm64)
- Poppler

## Setup

1. Install manual dependencies as mentioned above.

2. Set up your API keys and endpoints in the secret folder:
   - Create a new file `secret/openai_api_key` and paste your OpenAI API key into it.
   - Create a new file `secret/pinecone_api_key` and paste your Pinecone API key into it.
   - When setting up your Pinecone index, use a vector size of 1536 and keep all the default settings the same.
   - Create a new file `secret/pinecone_api_endpoint` and paste your Pinecone API endpoint into it.

3. Running the development environment:
   - Install JavaScript package dependencies: `npm install`
   - Run the Golang web server (default port: 8100): `npm start`
   - In another terminal window, run Webpack to compile the JS code and create a `bundle.js` file: `npm run dev`
   - Visit the local version of the site at [http://localhost:8100](http://localhost:8100)

## Screenshots

In the example screenshots, I uploaded a couple of books by Plato and some letters by Alexander Hamilton, showcasing the ability of AImate to answer questions based on the uploaded content.

### Uploading Files
![Uploading Files](upload_screenshots.png)

### Asking Questions
![Asking Questions](question_screenshots.png)

## Under the Hood

The Golang server uses POST APIs to process incoming uploads and respond to questions:

- `/upload` for uploading files
- `/api/question` for answering questions

All API endpoints are declared in the `vault-web-server/main.go` file.

### Uploading Files and Processing Them into Embeddings

The `vault-web-server/postapi/fileupload.go` file contains the `UploadHandler` logic for handling incoming uploads on the backend. The `UploadHandler` function is responsible for handling file uploads (with a maximum total upload size of 300 MB) and processing them into embeddings to store in Pinecone. It accepts PDF, epub, .docx, and plain text files, extracts text from them, and divides the content into chunks. Using the OpenAI API, it obtains embeddings for each chunk and upserts (inserts or updates) the embeddings into Pinecone. The function returns a JSON response containing information about the uploaded files and their processing status.

### Storing Embeddings into Pinecone DB

After getting OpenAI embeddings for each chunk of an uploaded file, the server stores all of the embeddings, along with associated metadata, in the Pinecone DB. The metadata includes file_name, start, end, title, and text. This metadata is useful for providing context to the embeddings and is used to display additional information about the matched embeddings when retrieving results from the Pinecone database.

### Answering Questions

The `QuestionHandler` function in `vault-web-server/postapi/questions.go` is responsible for handling all incoming questions. When a question is entered on the frontend and the user presses "search" (or enter), the server uses the OpenAI embeddings API once again to get an embedding for the question (query vector). This query vector is used to query the Pinecone DB to get the most relevant context for the question. Finally, a prompt is built by packing the most relevant context + the question in a prompt string that adheres to OpenAI token limits.

## Frontend Information

The frontend of AImate is built using React.js and Less for styling.

## Generative Question-Answering with Long-Term Memory

If you'd like to read more about this topic, I recommend this post from the Pinecone blog: [Generative Question-Answering with Long-Term Memory](https://www.pinecone.io/learn/openai-gen-qa/).

I hope you enjoy using AImate! If you have any questions or feedback, feel free to reach out!
