BUILD_DIR := build

.PHONY: all clean compile deb

all: deb

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

compile: $(BUILD_DIR)
	cd $(BUILD_DIR) && cmake ..
	cd $(BUILD_DIR) && $(MAKE)

deb: compile
	cd $(BUILD_DIR) && cpack -G DEB
	mv $(BUILD_DIR)/*.deb .

clean:
	rm -rf $(BUILD_DIR)
	rm -f *.deb
