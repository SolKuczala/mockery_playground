# mockery_playground

Is used for structs with methods
This case is used with a pkg appart. For mocks from main `--inpackage` flag needs
to be provided in order to avoid import issues

1. Add an interface with methods from that func
2. Run `mockery --all`
3. A new folder `mocks` is created to be used :party:
4. Mockery funcs can be call with .On and .Return, passing the correspondent
   parameters
