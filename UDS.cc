#include<uv.h>
#include<stdlib.h>
#include<iostream>
#include<unistd.h>
#include<string.h>
#include <stdio.h>

using namespace std;

std::string addr="localhost";
uv_loop_t main_loop;
uv_pipe_t uds_sock;

uv_connect_t conn;
uv_stream_t *conn_handle;
struct sockaddr_in server_sock;

void OnWrite(uv_write_t *req, int status) {
    return;
}

void WriteMessage(uv_stream_t *stream,int send_back){
    uv_write_t *req = new (uv_write_t);
    std::string response_buf("10");
    uv_buf_t req1= uv_buf_init((char *)response_buf.c_str(), response_buf.length());
    uv_write((uv_write_t *)req, stream, &req1, 1, [](uv_write_t *req, int status) { OnWrite(req, status); });
}

static void alloc_buffer(uv_handle_t *handle, size_t suggested_size, uv_buf_t *buf) {
         *buf = uv_buf_init((char*) malloc(suggested_size), suggested_size);
}

string Filename(const char *Message,ssize_t nread){
    std::string buffer;
    for(int i=0;i<nread;i++){
        buffer+=Message[i];
    }
    return buffer;
}

void OnRead(uv_stream_t *req,ssize_t nread, const uv_buf_t *buf){
        if (nread > 0) {
        WriteMessage(req,10);
    }
}

void OnConnect(uv_connect_t *conn,int status){
    if(status==0){
        uv_read_start(conn->handle, alloc_buffer,  [](uv_stream_t *stream, ssize_t nread, const uv_buf_t *buf) { OnRead(stream, nread, buf);  });
        conn_handle = conn->handle;
    }
}

int main(int argc,char** argv){
	uv_loop_init(&main_loop);
	uv_pipe_init(&main_loop, &uds_sock, 0);
	std::string uds_sock_path=argv[1];
	uv_pipe_connect(&conn, &uds_sock, uds_sock_path.c_str(), [](uv_connect_t *conn, int status) { OnConnect(conn, status);});
	uv_run(&main_loop, UV_RUN_DEFAULT);
}
