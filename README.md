# Service Repository Setup Guide

This guide will help you set up the template service repository as quickly as possible.

## Steps

1. Navigate to the `pkg/` directory and clone the proto repository. You can do this using HTTPS:
    ```
    git clone https://github.com/Dial-Afrika/proto.git
    ```
    Or using SSH:
    ```
    git clone git@github.com:Dial-Afrika/proto.git
    ```

2. Navigate into the `proto/` directory and generate the service stubs for Go by running:
    ```
    make proto
    ```

3. Navigate back to the root directory of the project.

4. In `pkg/auth/auth_inteceptors.go` and `pkg/auth/auth.go`, uncomment the interceptor and main auth method.

5. Implement the server configurations in `cmd/server/main.go`.

6. Add your gRPC services in the `pkg/services` directory. Remember to name them accordingly in one package name.

7. Set up the database either in the `docker-compose.yml` file or use a development environment database. If you choose to use a development environment database, you can request us to set it up with the development and staging environment pipelines.

## Conclusion

By following these steps, you should be able to quickly set up and start using the template service repository. If you encounter any issues, please raise an issue in the repository.
