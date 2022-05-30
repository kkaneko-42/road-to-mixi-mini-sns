NAME := road_to_mixi_1
PORT := 1323

$(NAME):
	docker-compose up -d

all: $(NAME)

clean:
	docker-compose down

re: clean all
