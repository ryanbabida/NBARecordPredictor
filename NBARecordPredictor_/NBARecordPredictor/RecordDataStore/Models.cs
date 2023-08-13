namespace NBARecordPredictor.RecordDataStore
{
    public class Record
    {
        public int GamesPlayed { get; set; }
        public int Wins { get; set; }
        public int Losses { get; set; }
        public decimal WinPercentage { get; set; }
        public decimal Minutes { get; set; }
        public decimal Points { get; set; }
        public decimal FieldGoalsMade { get; set; }
        public decimal FieldGoalsAttempted { get; set; }
        public decimal FieldGoalPercentage { get; set; }
        public decimal ThreesMade { get; set; }
        public decimal ThreesAttempted { get; set; }
        public decimal ThreePercentage { get; set; }
        public decimal FreeThrowsMade { get; set; }
        public decimal FreeThrowsAttempted { get; set; }
        public decimal FreeThrowPercentage { get; set; }
        public decimal OffensiveRebounds { get; set; }
        public decimal DefensiveRebounds { get; set; }
        public decimal Rebounds { get; set; }
        public decimal Assists { get; set; }
        public decimal Turnovers { get; set; }
        public decimal Steals { get; set; }
        public decimal Blocks { get; set; }
        public decimal BlocksAgainst { get; set; }
        public decimal PersonalFouls { get; set; }
        public decimal PersonalFoulsAgainst { get; set; }
        public decimal PlusMinus { get; set; }
    }

    public class RecordDataSet
    {
        public List<List<decimal>> FeatureSet { get; set; }
        public List<decimal> TargetSet { get; set; }

        public RecordDataSet()
        {
            FeatureSet = new List<List<decimal>>();
            TargetSet = new List<decimal>();
        }
    }
}
