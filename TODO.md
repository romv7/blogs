TODO
====
- Plan out a logic to where which we shall execute an operation, there are only two possible kinds of operation,
  via **Store** or a **gRPC endpoint**.

  The algorithm that we shall follow is described below:

  1. If there is an incoming gQL query, decide where we shall fetch or mutate the data.
  2. If there is an online gRPC endpoint we will use that gRPC connection for operation,
     instead, we will be using the **store.**
  3. If the type of operation is a query we will use the data loader to avoid database requests
     bottleneck if however the operation is a mutation we'll proceed calling just by calling
     the resolver associated with the operation.
  4. Done.