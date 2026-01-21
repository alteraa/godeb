cmake_minimum_required(VERSION 3.10)
project(example-deb VERSION 1.0.0)

add_executable(example-deb-bin main.cpp)

install(TARGETS example-deb-bin DESTINATION bin)
install(FILES example-deb.service DESTINATION /etc/systemd/system)



set(CPACK_PACKAGE_NAME "example-deb-pkg")
set(CPACK_PACKAGE_VERSION "1.0.0")
set(CPACK_PACKAGE_DESCRIPTION "A sample package for testing purposes")
set(CPACK_PACKAGE_CONTACT "developer@example.com")
set(CPACK_GENERATOR "DEB")
set(CPACK_PACKAGING_INSTALL_PREFIX "/usr")


set(CPACK_DEBIAN_PACKAGE_CONTROL_EXTRA
    "${CMAKE_CURRENT_SOURCE_DIR}/scripts/preinst"
    "${CMAKE_CURRENT_SOURCE_DIR}/scripts/postinst"
    "${CMAKE_CURRENT_SOURCE_DIR}/scripts/prerm"
    "${CMAKE_CURRENT_SOURCE_DIR}/scripts/postrm"
)

include(CPack)
