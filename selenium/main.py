import time

from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait

from selenium import webdriver


def test_create_comment():
    driver = webdriver.Chrome()
    driver.get("https://app.glyph.local/article/4656362c-d9ec-4005-bb38-daf07f004e26")
    time.sleep(5)  # debug

    author_input = driver.find_element(By.NAME, "author-name")
    author_input.send_keys("Test Reader")

    content_input = driver.find_element(By.NAME, "content")
    content_input.send_keys("This is a test comment.")

    submit_button = driver.find_element(By.ID, "submit-comment")
    submit_button.click()

    driver.get("https://app.glyph.local/admin/comments")

    time.sleep(1)
    no_comments = driver.find_elements(By.XPATH, "//*[contains(text(), 'No pending comments')]")
    assert not no_comments

    driver.quit()


def test_approve_comment():
    driver = webdriver.Chrome()
    driver.get("https://app.glyph.local/admin/comments")

    wait = WebDriverWait(driver, 5)

    wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "[data-testid='comments-section']")))

    comments = driver.find_elements(By.CSS_SELECTOR, "[data-testid='admin-comment']")

    target_comment = None
    for c in comments:
        content = c.find_element(By.CSS_SELECTOR, "[data-testid='comment-content']").text
        if "This is a test comment." in content:
            target_comment = c
            break

    assert target_comment is not None, "Test comment not found"

    approve_button = target_comment.find_element(By.CSS_SELECTOR, "[data-testid='approve-btn']")
    approve_button.click()

    wait.until(EC.text_to_be_present_in_element((By.TAG_NAME, "body"), "No pending comments 🎉"))

    driver.quit()


def test_delete_comment():
    driver = webdriver.Chrome()
    driver.get("https://app.glyph.local/admin/comments")

    wait = WebDriverWait(driver, 5)

    wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "[data-testid='comments-section']")))

    comments = driver.find_elements(By.CSS_SELECTOR, "[data-testid='admin-comment']")

    target_comment = None
    for c in comments:
        content = c.find_element(By.CSS_SELECTOR, "[data-testid='comment-content']").text
        if "This is a test comment." in content:
            target_comment = c
            break

    assert target_comment is not None, "Test comment not found"

    delete_button = target_comment.find_element(By.CSS_SELECTOR, "[data-testid='delete-btn']")
    delete_button.click()

    wait.until(EC.text_to_be_present_in_element((By.TAG_NAME, "body"), "No pending comments 🎉"))

    driver.get("https://app.glyph.local/article/4656362c-d9ec-4005-bb38-daf07f004e26")
    time.sleep(1)

    wait.until(EC.text_to_be_present_in_element((By.TAG_NAME, "body"), "No comments yet."))

    driver.quit()


if __name__ == "__main__":
    # test_create_comment()
    # test_approve_comment()

    test_create_comment()
    test_delete_comment()
