using Microsoft.AspNetCore.Mvc;
using NBARecordPredictor.RecordDataStore;

namespace NBARecordPredictor.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class RecordsController : ControllerBase
    {
        private readonly ILogger<RecordsController> _logger;
        private readonly IRecordDataStore _recordDataStore;

        public RecordsController(ILogger<RecordsController> logger, IRecordDataStore recordDataStore)
        {
            _logger = logger;
            _recordDataStore = recordDataStore;
        }

        public IEnumerable<Record> Get()
        {
            return _recordDataStore.GetAll();
        }

        [HttpGet("data")]
        public RecordDataSet GetDataSet()
        {
            return _recordDataStore.GetDataSet();
        }
    }
}