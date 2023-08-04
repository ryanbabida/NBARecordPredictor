using Microsoft.Extensions.Configuration;
using NBARecordPredictor.RecordDataStore;

namespace NBARecordPredictor.Test
{
    [TestClass]
    public class CsvDataStoreTest
    {
        private IConfiguration _testConfiguration { get; set; }

        public CsvDataStoreTest()
        {
            _testConfiguration = new ConfigurationBuilder()
                .AddJsonFile("appsettings.json")
                .Build();
        }

        [TestMethod]
        public void GetAllTest()
        {
            var expected = new Record()
            {
                GamesPlayed = 82,
                Wins = 15,
                Losses = 67,
                WinPercentage = 0.183m,
                Minutes = 50.1m,
                Points = 100.6m,
                FieldGoalsMade = 37.4m,
                FieldGoalsAttempted = 85.0m,
                FieldGoalPercentage = 44.0m,
                ThreesMade = 5.7m,
                ThreesAttempted = 16.2m,
                ThreePercentage = 35.1m,
                FreeThrowsMade = 20.1m,
                FreeThrowsAttempted = 26.8m,
                FreeThrowPercentage = 75.0m,
                OffensiveRebounds = 13.3m,
                DefensiveRebounds = 26.7m,
                Rebounds = 40.0m,
                Assists = 21.9m,
                Turnovers = 16.4m,
                Steals = 9.9m,
                Blocks = 3.8m,
                BlocksAgainst = 6.8m,
                PersonalFouls = 21.7m,
                PersonalFoulsAgainst = 0.1m,
                PlusMinus = -7.3m
            };

            var datastore = new RecordCsvDataStore(_testConfiguration);
            var results = datastore.GetAll();
            Assert.AreEqual(1, results.Count);

            var actual = results[0];
            AreRecordsEqual(expected, actual);
        }

        private void AreRecordsEqual(Record expected, Record actual)
        {
            Assert.AreEqual(expected.GamesPlayed, actual.GamesPlayed);
            Assert.AreEqual(expected.Wins, actual.Wins);
            Assert.AreEqual(expected.Losses, actual.Losses);
            Assert.AreEqual(expected.WinPercentage, actual.WinPercentage);
            Assert.AreEqual(expected.Minutes, actual.Minutes);
            Assert.AreEqual(expected.Points, actual.Points);
            Assert.AreEqual(expected.FieldGoalsMade, actual.FieldGoalsMade);
            Assert.AreEqual(expected.FieldGoalsAttempted, actual.FieldGoalsAttempted);
            Assert.AreEqual(expected.FieldGoalPercentage, actual.FieldGoalPercentage);
            Assert.AreEqual(expected.ThreesMade, actual.ThreesMade);
            Assert.AreEqual(expected.ThreesAttempted, actual.ThreesAttempted);
            Assert.AreEqual(expected.ThreePercentage, actual.ThreePercentage);
            Assert.AreEqual(expected.FreeThrowsMade, actual.FreeThrowsMade);
            Assert.AreEqual(expected.FreeThrowsAttempted, actual.FreeThrowsAttempted);
            Assert.AreEqual(expected.FreeThrowPercentage, actual.FreeThrowPercentage);
            Assert.AreEqual(expected.OffensiveRebounds, actual.OffensiveRebounds);
            Assert.AreEqual(expected.DefensiveRebounds, actual.DefensiveRebounds);
            Assert.AreEqual(expected.Rebounds, actual.Rebounds);
            Assert.AreEqual(expected.Assists, actual.Assists);
            Assert.AreEqual(expected.Turnovers, actual.Turnovers);
            Assert.AreEqual(expected.Steals, actual.Steals);
            Assert.AreEqual(expected.Blocks, actual.Blocks);
            Assert.AreEqual(expected.BlocksAgainst, actual.BlocksAgainst);
            Assert.AreEqual(expected.PersonalFouls, actual.PersonalFouls);
            Assert.AreEqual(expected.PersonalFoulsAgainst, actual.PersonalFoulsAgainst);
            Assert.AreEqual(expected.PlusMinus, actual.PlusMinus);
        }
    }
}