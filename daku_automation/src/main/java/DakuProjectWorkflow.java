//package com.example.tests;

import java.awt.*;
import java.awt.datatransfer.StringSelection;
import java.awt.event.KeyEvent;
import java.util.ArrayList;
import java.util.concurrent.TimeUnit;

import org.apache.poi.ss.formula.functions.T;
import org.junit.BeforeClass;
import org.openqa.selenium.chrome.ChromeOptions;
import org.openqa.selenium.interactions.Actions;
import org.testng.annotations.*;
import static org.testng.Assert.*;
import org.openqa.selenium.*;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.support.ui.Select;

public class DakuProjectWorkflow {



    private WebDriver driver;
    private String baseUrl;
    private boolean acceptNextAlert = true;
    private StringBuffer verificationErrors = new StringBuffer();

    @BeforeTest()
    public void setUp() throws Exception {
        String my_driver_path = "/home/kushalj/Documents/daku-cloud-project/daku_automation/chromedriver";
        System.setProperty("webdriver.chrome.driver", my_driver_path);



        ChromeOptions options = new ChromeOptions();
        options.addArguments("--headless", "--disable-gpu", "--window-size=1920,1200","--ignore-certificate-errors");
        driver = new ChromeDriver(options);
//        baseUrl = "https://www.blazedemo.com/";
        driver.manage().timeouts().implicitlyWait(30, TimeUnit.SECONDS);
    }

    @Test(invocationCount = 1, threadPoolSize = 1)
    public void testDakuProjectWorkflow() throws Exception {

        System.out.println(Thread.currentThread().getId());
        // Label: Test
        // ERROR: Caught exception [ERROR: Unsupported command [resizeWindow | 1536,656 | ]]


        int n=50;



        ArrayList<String> emails=new ArrayList<String>(n);
        ArrayList<String> passwords=new ArrayList<String>(n);

        for(int i=0;i<n;i++)
        {
//            System.out.println(i);
            emails.add("email_conf3_"+Integer.toString(i)+"@gmail.com");
            passwords.add("password"+Integer.toString(i));
        }

        for(int i=0;i<n;i++)
        {
//            System.out.println(i);
            System.out.println(emails.get(i));
            System.out.println(passwords.get(i));
        }

        driver.get("http://192.168.32.220:5000");


        driver.manage().window().maximize();
        // Label: register

//        for(int i=0;i<n;i++)
//        {
//
//        }

        for(int i=0;i<n;i++) {
            System.out.println(i);
            driver.findElement(By.linkText("Register")).click();
//        driver.findElement(By.id("username")).click();
//        driver.findElement(By.id("username")).clear();
//        driver.findElement(By.id("username")).sendKeys("kushal19057");
//        driver.findElement(By.cssSelector("fieldset.form-group")).click();
            driver.findElement(By.id("email")).click();
            driver.findElement(By.id("email")).clear();
            driver.findElement(By.id("email")).sendKeys(emails.get(i));
            driver.findElement(By.id("password")).click();
            driver.findElement(By.id("password")).sendKeys(passwords.get(i));
            driver.findElement(By.id("confirm_password")).click();
            driver.findElement(By.id("confirm_password")).sendKeys(passwords.get(i));
            driver.findElement(By.id("submit")).click();
            // Label: login

            driver.findElement(By.linkText("Login")).click();
            Thread.sleep(1000);

            driver.findElement(By.id("email")).click();
            driver.findElement(By.id("email")).clear();
            driver.findElement(By.id("email")).sendKeys(emails.get(i));
            driver.findElement(By.id("password")).click();
            driver.findElement(By.id("password")).sendKeys(passwords.get(i));


            // ERROR: Caught exception [unknown command [typeSecret]]
            driver.findElement(By.id("submit")).click();
            // Label: upload_file
            driver.findElement(By.id("navbarDropdown")).click();
            driver.findElement(By.linkText("File Upload")).click();

            WebElement open_file_button = driver.findElement(By.id("fileupload"));
//
            JavascriptExecutor executor = (JavascriptExecutor) driver;


            Thread.sleep(1000);
//        driver.findElement(By.id("fileupload")).clear();
            driver.findElement(By.id("fileupload")).sendKeys("/home/kushalj/Downloads/file1.jpg");


            driver.findElement(By.xpath("//button[@onclick='uploadFile()']")).click();
            // Label: view_file
            Thread.sleep(5000);

            driver.findElement(By.id("fileupload")).sendKeys("/home/kushalj/Downloads/file2.mp3");

            driver.findElement(By.xpath("//button[@onclick='uploadFile()']")).click();
            // Label: view_file
            Thread.sleep(5000);


            driver.findElement(By.id("navbarDropdown")).click();
            driver.findElement(By.linkText("My Files")).click();

            Thread.sleep(1000);


            Thread.sleep(2000);
            driver.findElement(By.id("navbarDropdown")).click();
            driver.findElement(By.linkText("Run Beast")).click();

            Thread.sleep(1000);


            driver.findElement(By.linkText("Logout")).click();
            Thread.sleep(1000);
        }

    }

    @AfterTest(alwaysRun = true)
    public void tearDown() throws Exception {
//        driver.quit();
        String verificationErrorString = verificationErrors.toString();
        if (!"".equals(verificationErrorString)) {
            fail(verificationErrorString);
        }
    }

    private boolean isElementPresent(By by) {
        try {
            driver.findElement(by);
            return true;
        } catch (NoSuchElementException e) {
            return false;
        }
    }

    private boolean isAlertPresent() {
        try {
            driver.switchTo().alert();
            return true;
        } catch (NoAlertPresentException e) {
            return false;
        }
    }

    private String closeAlertAndGetItsText() {
        try {
            Alert alert = driver.switchTo().alert();
            String alertText = alert.getText();
            if (acceptNextAlert) {
                alert.accept();
            } else {
                alert.dismiss();
            }
            return alertText;
        } finally {
            acceptNextAlert = true;
        }
    }
}
