package tabula.database.engine.model;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import tabula.database.query.QueryUtils;
import tabula.database.table.Header;
import tabula.database.table.Tabula;
import tabula.logger.LoggerUtil;

@AllArgsConstructor
@Getter
public class ProtoSQL {

    private Engine engine;
    private String connStr;
    private String dbName;

    @Setter
    private String queries;

    private int borderStyle;
    private String destFolder;
    private String headerStyleLink;

    public enum Option {
        RUN(1), TABLES(2), TABLE_INFO(3), PING(4);

        private int id;

        Option(int id) {
            this.id = id;
        }

        public int getId() {
            return this.id;
        }

        public static Option get(int id) {
            for (Option o : Option.values()) {
                if (o.id == id) {
                    return o;
                }
            }
            return null;
        }
    }

    public enum Engine {
        INFORMIX("com.informix.jdbc.IfxDriver");

        private String driver;

        Engine(String driver) {
            this.driver = driver;
        }

        public String getDriver() {
            return this.driver;
        }
    }

    public Connection getConnection() {
        try {
            Class.forName(this.engine.getDriver());
            return DriverManager.getConnection(this.connStr);
        } catch (ClassNotFoundException | SQLException e) {
            System.out.println("[ERROR] connecting DB: ".formatted(e.getMessage()));
            return null;
        }
    }

    public void run() {
        try (Connection connection = getConnection()) {
            if (QueryUtils.isSelectQuery(queries)) {
                executeSelect(connection);
            } else {
                execute(connection);
            }
        } catch (SQLException e) {
            System.err.printf("[ERROR] %s".formatted(e.getMessage()));
        }
    }

    public void execute(Connection db) {
        if (!QueryUtils.containsSemicolonInMiddle(queries)) {
            try (PreparedStatement stmt = db.prepareStatement(queries)) {
                var rowsAffected = stmt.executeUpdate();

                if (QueryUtils.isInsertUpdateOrDelete(queries)) {
                    System.out.printf("  Row(s) affected: %d", rowsAffected);
                } else {
                    System.out.println("  Statement executed correctly.");
                }
            } catch (SQLException e) {
                LoggerUtil.errorf("Error executing query: %s", e.getMessage());
                System.out.printf("[ERROR] %s", e.getMessage());
            }
        } else {
            var queryList = QueryUtils.splitQueries(queries);
            var results = new String[queryList.size()];

            for (int i = 0; i < queryList.size(); i++) {
                var q = queryList.get(i);
                try (PreparedStatement stmt = db.prepareStatement(q)) {
                    var rowsAffected = stmt.executeUpdate();

                    if (QueryUtils.isInsertUpdateOrDelete(q)) {
                        results[i] = String.format("%d)   Row(s) affected: %d\n", i + 1, rowsAffected);
                    } else {
                        results[i] = String.format("%d)   Statement executed correctly.\n", i + 1);
                    }
                } catch (SQLException e) {
                    LoggerUtil.errorf("Error executing query: %s", e.getMessage());
                    results[i] = String.format("%d)   %s\n", i + 1, e.getMessage());
                }
            }

            String filePath = Tabula.createTabulaFileFormat(destFolder);
            System.out.println("syn match tabulaStmtErr ' ' | hi link tabulaStmtErr ErrorMsg");
            System.out.println(filePath);

            Tabula.writeToFile(filePath, results);
        }
    }

    public void executeSelect(Connection connection) {
        try (Statement stmt = connection.createStatement(); ResultSet rs = stmt.executeQuery(queries)) {
            List<String> columns = new ArrayList<>();
            int columnCount = rs.getMetaData().getColumnCount();

            for (int i = 1; i <= columnCount; i++) {
                columns.add(rs.getMetaData().getColumnName(i));
            }

            Map<Integer, Header> headers = new HashMap<>();
            for (int i = 0; i < columns.size(); i++) {
                String name = " 󰠵 " + columns.get(i).toUpperCase();
                headers.put(i + 1, new Header(name, QueryUtils.unicodeLength(name) + 1));
            }

            List<List<String>> rows = new ArrayList<>();
            while (rs.next()) {
                List<String> resultRow = new ArrayList<>(columnCount);
                for (int i = 1; i <= columnCount; i++) {
                    String value = rs.getString(i);
                    if (value == null) {
                        value = "NULL";
                    }
                    resultRow.add(" " + value);
                    int valueLength = QueryUtils.unicodeLength(value) + 2;
                    if (headers.get(i).getLength() < valueLength) {
                        headers.put(i, new Header(headers.get(i).getName(), valueLength));
                    }
                }
                rows.add(resultRow);
            }

            if (!rows.isEmpty()) {
                System.out.println(headers.toString());
                System.out.println(rows.toString());
                var tabula = new Tabula(destFolder, borderStyle, headerStyleLink, headers, rows);
                tabula.generate();
            } else {
                System.out.println("  Query has returned 0 results.");
            }
        } catch (SQLException e) {
            System.out.printf("[ERROR] SQL %s", e.getMessage());
        }
    }

    public void getTables() {
        List<String> values = new ArrayList<>();

        try (Connection connection = getConnection();
                Statement stmt = connection.createStatement();
                ResultSet rs = stmt.executeQuery(this.queries)) {

            while (rs.next()) {
                String table = rs.getString(1);
                values.add(table);
            }
        } catch (SQLException e) {
            System.out.println("[ERROR] %s".formatted(e.getMessage()));
        }
        System.out.println(values);
    }
}
