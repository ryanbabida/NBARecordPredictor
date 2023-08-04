using NBARecordPredictor.RecordDataStore;
using OpenTelemetry.Trace;

namespace NBARecordPredictor
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            // Add services to the container.
            builder.Services.AddSingleton<IRecordDataStore, RecordCsvDataStore>();

            builder.Services.AddControllers();
            // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
            builder.Services.AddEndpointsApiExplorer();
            builder.Services.AddSwaggerGen();

            builder.Services.AddOpenTelemetry()
                .WithTracing(
                    builder =>
                        builder
                            .AddSource("NBARecordPredictor")
                            .AddAspNetCoreInstrumentation(options =>
                                {
                                    options.RecordException = true;
                                })
                            .AddHttpClientInstrumentation()
                            .AddSqlClientInstrumentation(options =>
                                {
                                    options.SetDbStatementForText = true;
                                    options.SetDbStatementForStoredProcedure = true;
                                    options.RecordException = true;
                                })
                            .AddConsoleExporter()
                );

            var app = builder.Build();

            // Configure the HTTP request pipeline.
            if (app.Environment.IsDevelopment())
            {
                app.UseSwagger();
                app.UseSwaggerUI();
            }

            app.UseHttpsRedirection();

            app.UseAuthorization();


            app.MapControllers();

            app.Run();
        }
    }
}