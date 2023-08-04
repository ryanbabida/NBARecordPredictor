using System.Globalization;
using System.Text;
using CsvHelper;
using CsvHelper.Configuration;

namespace NBARecordPredictor.RecordDataStore
{
    public class RecordCsvDataStore : IRecordDataStore
    {
        private List<Record> _records { get; }

        public RecordCsvDataStore(IConfiguration config)
        {
            _records = new List<Record>();
            var filename = config["RecordCsvDataStoreConfig:Filename"];
            var csvConfig = new CsvConfiguration(CultureInfo.InvariantCulture)
            {
                Encoding = Encoding.UTF8,
                Delimiter = ",",
                HasHeaderRecord = false
            };

            using (var fs = File.Open(filename, FileMode.Open, FileAccess.Read, FileShare.Read))
            {
                using (var textReader = new StreamReader(fs, Encoding.UTF8))
                using (var csvReader = new CsvReader(textReader, csvConfig))
                {
                    csvReader.Context.RegisterClassMap<RecordMap>();
                    var data = csvReader.GetRecords<Record>();

                    foreach (var record in data)
                    {
                        _records.Add(record);
                    }
                }
            }
        }

        public List<Record> GetAll()
        {
            return _records;
        }
    }

    public class RecordMap : ClassMap<Record>
    {
        public RecordMap()
        {
            Map(r => r.GamesPlayed).Index(0);
            Map(r => r.Wins).Index(1);
            Map(r => r.Losses).Index(2);
            Map(r => r.WinPercentage).Index(3);
            Map(r => r.Minutes).Index(4);
            Map(r => r.Points).Index(5);
            Map(r => r.FieldGoalsMade).Index(6);
            Map(r => r.FieldGoalsAttempted).Index(7);
            Map(r => r.FieldGoalPercentage).Index(8);
            Map(r => r.ThreesMade).Index(9);
            Map(r => r.ThreesAttempted).Index(10);
            Map(r => r.ThreePercentage).Index(11);
            Map(r => r.FreeThrowsMade).Index(12);
            Map(r => r.FreeThrowsAttempted).Index(13);
            Map(r => r.FreeThrowPercentage).Index(14);
            Map(r => r.OffensiveRebounds).Index(15);
            Map(r => r.DefensiveRebounds).Index(16);
            Map(r => r.Rebounds).Index(17);
            Map(r => r.Assists).Index(18);
            Map(r => r.Turnovers).Index(19);
            Map(r => r.Steals).Index(20);
            Map(r => r.Blocks).Index(21);
            Map(r => r.BlocksAgainst).Index(22);
            Map(r => r.PersonalFouls).Index(23);
            Map(r => r.PersonalFoulsAgainst).Index(24);
            Map(r => r.PlusMinus).Index(25);
        }
    }

    public class RecordCsvDataStoreConfig
    {
        public static string SectionName = "RecordCsvDataStoreConfig";
        public string? Filename { get; }
    }
}
