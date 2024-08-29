## nvim-tabula
*Minimal Multi database client for Neovim*

## Caveats
- These dependencies are required to be installed: `Go`. 
- Java is not required. Only if you want to use databases which require Java, in that case `Java 21 or newer` and `Maven` are required.
- For the sake of simplicity, **this plugin is STATELESS**. It does not use database sessions or keep states after Neovim is closed.
- This plugin has been developed on and for `Linux` following open source philosophy.

## Supported Databases
##### Databases not marked will be supported in the future
- Go implemented
    - [x] MongoDB
    - [x] MySQL
    - [x] MS-SQL (Microsoft)
    - [ ] Neo4j
    - [ ] Oracle
    - [x] PostgreSQL
    - [ ] Redis
    - [ ] SQLite
- Java implemented
    - [ ] IBM DB2
    - [x] IBM Informix
