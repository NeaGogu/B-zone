const { expect } = require("chai");
const puppeteer = require("puppeteer");

let browser;
let page;

describe("Login Test", () => {
  before(async () => {
    browser = await puppeteer.launch({
      headless: false,
      args: ["--start-maximized"],
    });

    page = await browser.newPage();
    await page.setExtraHTTPHeaders({
      "Accept-Language": "en-US,en;q=0.9",
    });
    await page.setViewport({ width: 1366, height: 768 });
  });

  after(async () => {});
  it("Should not be accessible home page before login", async () => {
    await page.goto("http://localhost:3000/home");
    await page.waitForNavigation();
    await page.waitForTimeout(1000);
    const url = await page.url().split("/").pop();
    expect(url).to.equal("login");
  }).timeout(10000);

  it("Should fill in the username field", async () => {
    await page.goto("http://localhost:3000/login");
    await page.waitForTimeout(3500);
    page.type("input[data-testid=username-input]", "sep@bumbal.eu");
    await page.waitForTimeout(1000);
    const value = await page.$eval(
      "input[data-testid=username-input]",
      (input) => input.value
    );
    await page.waitForTimeout(1000);
    expect(value).to.equal("sep@bumbal.eu");
  }).timeout(10000);

  it("Should fill in the password field", async () => {
    await page.waitForTimeout(1000);
    page.type("input[data-testid=password-input]", "cW$#Qbph0524");
    await page.waitForTimeout(1000);
    const value = await page.$eval(
      "input[data-testid=password-input]",
      (input) => input.value
    );
    expect(value).to.equal("cW$#Qbph0524");
  }).timeout(10000);

  it("Should click the login button", async () => {
    await page.waitForTimeout(1000);
    await page.click("button[data-testid=button]");
    await page.waitForNavigation();
    await page.waitForTimeout(2000);
  }).timeout(10000);

  it("Should redirect to home page after successful login", async () => {
    await page.waitForTimeout(1000);
    await page.waitForSelector("svg[data-icon=user]");
    const url = await page.url().split("/").pop();
    expect(url).to.equal("home");
  }).timeout(5000);
  it("Should not be accessible login page after login", async () => {
    await page.waitForTimeout(1000);
    await page.goto("http://localhost:3000/login");
    await page.waitForNavigation();
    await page.waitForTimeout(2000);
    const url = await page.url().split("/").pop();
    expect(url).to.equal("home");
  }).timeout(10000);
});

describe("Map Test", () => {
  before(async () => {});

  after(async () => {});

  it("Should map zoom in", async () => {
    await page.waitForSelector("main[id=map]");

    page.hover("main[id=map]");
    await page.mouse.wheel({ deltaY: -100 });
    await page.waitForTimeout(1500);
  }).timeout(10000);

  it("Should map zoom out", async () => {
    await page.waitForSelector("main[id=map]");

    page.hover("main[id=map]");
    await page.mouse.wheel({ deltaY: 100 });
    await page.waitForTimeout(1500);
  }).timeout(10000);

  it("Should map zoom in button click", async () => {
    await page.waitForSelector("a.leaflet-control-zoom-in");
    page.click("a.leaflet-control-zoom-in");
    await page.waitForTimeout(1000);
  }).timeout(10000);

  it("Should map zoom out button click", async () => {
    await page.waitForSelector("a.leaflet-control-zoom-in");
    page.click("a.leaflet-control-zoom-out");
    await page.waitForTimeout(1000);
  }).timeout(10000);

  it("Should add new marker", async () => {
    await page.waitForSelector("main[id=map]");
    //scroll map zoom
    page.hover("main[id=map]");
    await page.waitForTimeout(1000);
    await page.mouse.click(500, 500);
    await page.waitForTimeout(1000);
  }).timeout(10000);

  it("Should map control layers", async () => {
    await page.waitForSelector("a.leaflet-control-layers-toggle");
    page.hover("a.leaflet-control-layers-toggle");
    await page.waitForTimeout(1000);
    //select checkboxs
    await page.waitForSelector("input.leaflet-control-layers-selector");
    const checkboxs = await page.$$("input.leaflet-control-layers-selector");
    checkboxs[0].click();
    await page.waitForTimeout(1000);
    checkboxs[0].click();
    await page.waitForTimeout(1000);
    checkboxs[1].click();
    await page.waitForTimeout(1000);
    checkboxs[1].click();
    await page.waitForTimeout(1000);
  }).timeout(10000);
});

describe("Sidebar Test", () => {
  before(async () => {});
  after(async () => {});
  it("Expand menu items", async () => {
    await page.waitForSelector("div[role=menuitem]");
    const menuitems = await page.$$("div[role=menuitem]");
    for (let i = 0; i < menuitems.length; i++) {
      await page.waitForTimeout(1000);
      menuitems[i].click();
    }
    await page.waitForTimeout(1000);
  }).timeout(10000);
  it("Should radio button click", async () => {
    await page.waitForTimeout(1000);
    const radioButtons = await page.$$("label.ant-radio-button-wrapper");
    radioButtons[0].click();
    await page.waitForTimeout(1000);
    radioButtons[1].click();
    await page.waitForTimeout(1000);
    intensityInput = await page.$("input[data-testid='intensity-input']");
    intensityInput.click();
    await page.waitForTimeout(1000);
    for (let i = 0; i < 3; i++) {
      await intensityInput.press("Backspace");
    }
    await page.waitForTimeout(1000);
    intensityInput.type("500");
    radioButtons[0].click();
    await page.waitForTimeout(1000);
  }).timeout(20000);
  it("The button should be activated when the value is entered into the inputs", async () => {
    await page.waitForTimeout(1000);
    const isDisabled = (await page.$("button[disabled][type=submit]")) == null;
    expect(isDisabled).to.be.true;
  }).timeout(10000);
  it("Zones input should be working and button enable", async () => {
    const inputs = await page.$$("input.ant-input");
    await page.waitForTimeout(1000);
    inputs[0].click();
    for (let i = 0; i < 10; i++) {
      await inputs[0].press("Backspace");
    }
    await page.waitForTimeout(1000);
    inputs[0].type("3");
    await page.waitForTimeout(1000);
    inputs[1].click();
    for (let i = 0; i < 10; i++) {
      await inputs[1].press("Backspace");
    }
    await page.waitForTimeout(1000);
    inputs[1].type("3");
    await page.waitForTimeout(1000);
    const inputValue = await inputs[0]
      .getProperty("value")
      .then((v) => v.jsonValue());
    await page.waitForTimeout(1000);
    expect(inputValue).to.equal("3");
  }).timeout(30000);
  it("Should click the Options for Algorithms (The standard is closed when advanced is selected)", async () => {
    //sidebar scroll down
    await page.waitForSelector("form.ant-form");
    await page.hover("form.ant-form");
    await page.waitForTimeout(1000);
    await page.mouse.wheel({ deltaY: 100 });
    await page.waitForTimeout(1000);
    const radioButtons = await page.$$("input.ant-radio-input");
    await page.waitForTimeout(1000);
    radioButtons[0].click();
    await page.waitForTimeout(1000);
    radioButtons[3].click();
    await page.waitForTimeout(1000);
    radioButtons[2].click();
    await page.waitForTimeout(1000);
    radioButtons[1].click();
    await page.waitForTimeout(1000);
    radioButtons[0].click();
    await page.waitForTimeout(1000);
    const isDisabled = await page.evaluate(
      (element) => element.disabled,
      radioButtons[2]
    );
    expect(!isDisabled).to.be.true;
  }).timeout(30000);
  it("Should type Number of Zones input and click the button", async () => {
    await page.waitForSelector("input.ant-input");
    const inputs = await page.$$("input.ant-input");
    await page.waitForTimeout(1000);
    inputs[2].click();
    for (let i = 0; i < 10; i++) {
      await inputs[2].press("Backspace");
    }
    await page.waitForTimeout(1000);
    inputs[2].type("3");
    await page.waitForTimeout(1000);
    const buttons = await page.$$("button.ant-btn-primary");
    buttons[1].click();
    await page.waitForTimeout(5000);
  }).timeout(30000);

  it("Initial zone collapse expand buttons click", async () => {
    await page.waitForSelector("div.ant-collapse-header");
    const InitialZoneButtons = await page.$$("div.ant-collapse-header");
    await page.waitForTimeout(1000);
    for (let i = 0; i < InitialZoneButtons.length; i++) {
      await page.waitForTimeout(1000);
      InitialZoneButtons[i].click();
    }
  }).timeout(10000);
  it("Should save zone", async () => {
    page.on("dialog", async (dialog) => {
      await page.waitForTimeout(1000);
      await dialog.accept("test zone 4");
    });
    await page.waitForSelector("button.ant-btn-primary");
    await page.waitForTimeout(1000);
    const buttons = await page.$$("button.ant-btn-primary");
    buttons[0].click();
    await page.waitForTimeout(1000);

    await page.waitForTimeout(1000);
  }).timeout(10000);

  it("Should click view and compare button", async () => {
    await page.waitForSelector("button.ant-btn-default");
    buttons = await page.$$("button.ant-btn-default");
    await page.waitForTimeout(1000);
    buttons[1].click();
    await page.waitForTimeout(1000);
    buttons[2].click();
    await page.waitForTimeout(1000);
  }).timeout(10000);
});

describe("Logout Test", () => {
  before(async () => {});
  after(async () => {
    await browser.close();
  });
  it("Should logout success", async () => {
    await page.waitForSelector("svg[data-icon=user]");
    await page.waitForTimeout(1000);
    page.hover("svg[data-icon=user]");
    await page.waitForTimeout(1000);
    await page.waitForSelector("button[data-testid=btn]");
    page.click("button[data-testid=btn]");
    await page.waitForTimeout(4000);
  }).timeout(15000);
});
