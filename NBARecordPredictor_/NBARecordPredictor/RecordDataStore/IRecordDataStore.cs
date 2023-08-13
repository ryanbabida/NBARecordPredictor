namespace NBARecordPredictor.RecordDataStore
{
    public interface IRecordDataStore
    {
        public List<Record> GetAll();
        public RecordDataSet GetDataSet();
    }
}
