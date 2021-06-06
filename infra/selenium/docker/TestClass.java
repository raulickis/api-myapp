@Parameters({"Port"})
@BeforeClass
public void initiateDriver(String Port) throws MalformedURLException {
    if(Port.equalsIgnoreCase("5900"))
    {
        driver = new RemoteWebDriver(new URL("http:localhost:4444/wd/hub"), DesiredCapabilities.chrome());
        driver.manage().window().maximize();
    }
    // else if(Port.equalsIgnoreCase("9002")){
    //    driver = new RemoteWebDriver(new URL("http:localhost:4444/wd/hub"), DesiredCapabilities.firefox());
    //    driver.manage().window().maximize();
    // }
}
