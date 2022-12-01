
# Client server grpc example

Для запуска клиента и сервера нужно выполнить следующее:

    git clone https://github.com/Kolyan4ik99/grpc-client-server
    cd grpc-client-server

    make

    ./server

    ./client --user_name=Thomas --user_password=Anderson \
    --dial_interval=500ms --dial_deadline=11.24s \
    --buffer_size=6 --buffer_threshold=16s
