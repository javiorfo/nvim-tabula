package tabula;

import org.apache.commons.cli.DefaultParser;
import org.apache.commons.cli.Options;
import org.apache.commons.cli.ParseException;

import tabula.database.engine.model.ProtoSQL;
import tabula.database.factory.DBFactory;
import tabula.logger.LoggerUtil;

public class Main {
    public static void main(String[] args) {
        var options = new Options();
        options.addOption("e", "engine", true, "Database engine");
        options.addOption("c", "conn-str", true, "Database string connection");
        options.addOption("n", "dbname", true, "Database name");
        options.addOption("q", "queries", true, "Database queries semicolon-separated");
        options.addOption("b", "border-style", true, "Table border style");
        options.addOption("d", "dest-folder", true, "Destinated folder for tabula files");
        options.addOption("l", "tabula-log-file", true, "Neovim Tabula log file");
        options.addOption("o", "option", true, "Options to execute: 1:run/2:tables/3:table-info/4:ping");
        options.addOption("h", "header-style-link", true, "hi link header type");
        options.addOption("g", "log-debug", true, "Enable debug level logger");

        try {
            var cmd = new DefaultParser().parse(options, args);
            var proto = new ProtoSQL(
                    ProtoSQL.Engine.valueOf(cmd.getOptionValue("e").toUpperCase()),
                    cmd.getOptionValue("c"),
                    cmd.getOptionValue("n"),
                    cmd.getOptionValue("q"),
                    Integer.valueOf(cmd.getOptionValue("b", "1")),
                    cmd.getOptionValue("d", "/tmp"),
                    cmd.getOptionValue("h", "Type"));

            var op = ProtoSQL.Option.get(Integer.valueOf(cmd.getOptionValue("o", "1")));

            LoggerUtil.initialize(cmd.getOptionValue("l"), Boolean.valueOf(cmd.getOptionValue("g")));

            LoggerUtil.debug("ProtoSQL obj: %s".formatted(proto.toString()));

            DBFactory.context(op, proto);

            LoggerUtil.close();

        } catch (ParseException e) {
            System.out.println("[ERROR] parsing command line arguments: " + e.getMessage());
        }
    }
}
