import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SeeDenialExplanationPromptComponent } from './see-denial-explanation-prompt.component';

describe('SeeDenialExplanationPromptComponent', () => {
  let component: SeeDenialExplanationPromptComponent;
  let fixture: ComponentFixture<SeeDenialExplanationPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SeeDenialExplanationPromptComponent]
    });
    fixture = TestBed.createComponent(SeeDenialExplanationPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
