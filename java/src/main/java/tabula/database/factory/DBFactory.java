package tabula.database.factory;

import tabula.database.engine.Informix;
import tabula.database.engine.model.ProtoSQL;

public class DBFactory {

    public static void context(ProtoSQL.Option option, ProtoSQL proto) {
        switch (proto.getEngine()) {
            case INFORMIX -> run(new Informix(proto), option);
            default -> System.out.println("[ERROR] Engine does not exist");
        }
    }

    public static void run(Executor executor, ProtoSQL.Option option) {
        switch (option) {
            case RUN -> executor.run();
            case TABLES -> executor.getTables();
            case TABLE_INFO -> executor.getTableInfo();
            case PING -> executor.ping();
        }
    }
}
