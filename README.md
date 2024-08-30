# nvim-tabula
*Minimal Multi database client for Neovim*

## Caveats
- These dependencies are required to be installed: `Go`. 
- Java is not required. Only if you want to use databases which require Java, in that case `Java 21 or newer` and `Maven` are required.
- For the sake of simplicity, **this plugin is STATELESS**. It does not use database sessions or keep states after Neovim is closed.
- This plugin has been developed on and for `Linux` following open source philosophy.

## Supported Databases
#### Databases not marked will be supported in the future

| Database | Supported | Language required | NOTE |
| ------- | ------------- | ------ | ---- |
| MongoDB | :heavy_check_mark: | Go | Supported operations detailed here |
| MySQL | :heavy_check_mark: | Go | Supported operations detailed here |
| MS-SQL | :heavy_check_mark: | Go | Supported operations detailed here |
| PostgreSQL | :heavy_check_mark: | Go | Supported operations detailed here |
| Neo4j | :x: | Go | Future release |
| Oracle | :x: | Go | Future release |
| Redis | :x: | Go | Future release |
| SQLite | :x: | Go | Future release |
| IBM Informix | :heavy_check_mark: | Java | Supported operations detailed here |


## Installation
`Lazy`
```lua
{ 
    'javiorfo/nvim-tabula',
    dependencies = {
        'javiorfo/nvim-popcorn',
        'javiorfo/nvim-spinetta'
    },
    lazy = true,
    cmd = { "TabulaBuild" },
    ft = { "sql", "javascript" }, -- javascript if MongoDB is used
    build = function()
        -- Update the backend in every plugin change
        require'tabula.core'.build()
    end,
    opts = {
        -- This section is not required
        -- Only if you want to change default configurations
      
        -- Default keymaps
        commands = {
            -- Keymap in Normal mode to change DB with command :TabulaSelectDB
            select_db = '<C-space>',
            
            -- Keymap in Normal and Visual mode to execute a query
            execute = '<C-t>',
            
            -- Keymap in Normal mode to close all buffer results
            close = '<C-c>',
        },

        -- Command :TabulaShowDB
        view = {
            -- Show the user name
            show_user = true,
            
            -- Show the user password
            show_password = true,
        },

        -- Output buffer
        output = {
            -- Default dest folder where .tabula files are created
            dest_folder = "/tmp",

            -- Border style of the table result (1 to 6 to choose)
            border_style = 1,

            -- A "hi link column style" in header table results
            header_style_link = "Type",

            -- Height of the buffer table result
            buffer_height = 20,
        },

        -- Configuration of databases (host, port, credentials, etc)
        db = {
            -- Default DB when open a buffer
            default = 1,
        },

        -- For errors and debug purposes if anything goes wrong
        internal = {
            log_debug = false
        }
    }
}
```

## Configuration
#### Configure DB connections and credentials


## Usage


## Screenshots
### Simple use

<img src="https://github.com/javiorfo/img/blob/master/nvim-tabula/tabula.gif?raw=true" alt="nvim-tabula"/>

**NOTE:** The colorscheme **nox** from [nvim-nyctophilia](https://github.com/javiorfo/nvim-nyctophilia) is used in this image.

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
