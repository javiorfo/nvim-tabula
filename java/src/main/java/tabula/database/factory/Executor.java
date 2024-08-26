package tabula.database.factory;

public interface Executor {
    public void run();

    public void getTables();

    public void getTableInfo();

    public void ping();
}
